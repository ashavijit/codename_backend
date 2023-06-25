package middlewares

import (
	"codename_backend/database"
	"codename_backend/models"
	"codename_backend/utils"
	"context"
	// "fmt"
	"regexp"

	// "fmt"
	"time"

	// "encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	DB_NAME        = "codename"
	CollectionName = "users"
)
var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis-10351.c114.us-east-1-4.ec2.cloud.redislabs.com:10351",
		Password: "QIqrMPlHJJ8UuhFjr936khTiJwPE3ChP",
		DB:       0,
	})
}

func RegisterNewUsers(c *gin.Context) {
	collection := database.GetCollection(CollectionName)

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter := bson.M{"email": user.Email}
	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	err = rdb.Set(context.Background(), user.Email, user.Password, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Redis error"})
		return
	}

	// Insert the user into the database
	_, err = collection.InsertOne(context.Background(), &user)
	if err != nil {
		// Remove user from Redis if the database insertion fails
		rdb.Del(context.Background(), user.Email)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.Next()
}

func LoginMiddleware(c *gin.Context) {
	collection := database.GetCollection(CollectionName)
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		c.Abort()
		return
	}
	filter := bson.M{"email": user.Email, "password": user.Password}
	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		c.Abort()
		return
	}
	if count == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		c.Abort()
		return
	}

	c.Next()
}

func CheckLoggedIn(c *gin.Context) {
	collection := database.GetCollection(CollectionName)
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		c.Abort()
		return
	}
	filter := bson.M{"email": user.Email, "password": user.Password}
	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		c.Abort()
		return
	}
	if count == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		c.Abort()
		return
	}

	c.Next()
}

func ResetPassword(c *gin.Context) {
	collection := database.GetCollection(CollectionName)
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		c.Abort()
		return
	}
	filter := bson.M{"email": user.Email}
	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		c.Abort()
		return
	}
	if count == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		c.Abort()
		return
	}
	otp := utils.GenerateOTP()
	user.OTPTimestamp = time.Now()
	err = utils.SendEmail(user.Email, otp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		c.Abort()
		return
	}
	_, err = collection.UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"otp": otp, "otp_timestamp": user.OTPTimestamp}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		c.Abort()
		return
	}
	c.Next()
}

func VerifyPassword(c *gin.Context) {
	collection := database.GetCollection(CollectionName)
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	filter := bson.M{"email": user.Email, "otp": user.OTP}
	err := collection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Assuming OTPTimestamp is of type time.Time
	if time.Since(user.OTPTimestamp).Minutes() > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "OTP expired"})
		return
	}

	type ResetPasswordReqBody struct {
		Email       string `json:"email" binding:"required"`
		OTP         string `json:"otp" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}

	var resetPasswordReqBody ResetPasswordReqBody

	if err := c.ShouldBindJSON(&resetPasswordReqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if (utils.EmailValidate(resetPasswordReqBody.Email)){
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	if resetPasswordReqBody.Email != user.Email {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	if resetPasswordReqBody.OTP != user.OTP {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OTP"})
		return
	}

	// Update password
	update := bson.M{"$set": bson.M{"password": resetPasswordReqBody.NewPassword}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}


func ChangePassword(c *gin.Context) {
	collection := database.GetCollection(CollectionName)

	type PasswordReqBody struct {
		Email        string `json:"email" binding:"required"`
		OldPassword  string `json:"old_password" binding:"required"`
		NewPassword  string `json:"new_password" binding:"required"`
	}

	var passwordReqBody PasswordReqBody
	if err := c.ShouldBindJSON(&passwordReqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	EmailRegEx := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !EmailRegEx.MatchString(passwordReqBody.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	filter := bson.M{"email": passwordReqBody.Email}
	var user models.User
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if passwordReqBody.OldPassword != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	update := bson.M{"$set": bson.M{"password": passwordReqBody.NewPassword}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}

func UserNameChange(c *gin.Context) {
	collection := database.GetCollection(CollectionName)

	type UserNameReqBody struct {
		Email       string `json:"email" binding:"required"`
		NewUserName string `json:"new_username" binding:"required"`
		Password    string `json:"password" binding:"required"`
	}

	var userNameReqBody UserNameReqBody
	if err := c.ShouldBindJSON(&userNameReqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(userNameReqBody.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	findUser := bson.M{"email": userNameReqBody.Email, "password": userNameReqBody.Password}
	var user models.User
	err := collection.FindOne(context.Background(), findUser).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found or database error"})
		return
	}

	if user.Username == userNameReqBody.NewUserName {
		c.JSON(http.StatusBadRequest, gin.H{"error": "New username is the same as the current username"})
		return
	}

	updateUserName := bson.M{"$set": bson.M{"username": userNameReqBody.NewUserName}}
	_, err = collection.UpdateOne(context.Background(), findUser, updateUserName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update username"})
		return
	}

	c.Next()	
}

func UserEmailChange(c *gin.Context) {
	collection := database.GetCollection(CollectionName)

	type UserEmailReqBody struct {
		Email       string `json:"email" binding:"required"`
		NewEmail    string `json:"new_email" binding:"required"`
		Password    string `json:"password" binding:"required"`
	}

	var userEmailReqBody UserEmailReqBody
	if err := c.ShouldBindJSON(&userEmailReqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(userEmailReqBody.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	FindUser := bson.M{"email": userEmailReqBody.Email, "password": userEmailReqBody.Password}
	var user models.User
	err := collection.FindOne(context.Background(), FindUser).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found or database error"})
		return
	}
	// check for password match
	if user.Password != userEmailReqBody.Password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}

	if user.Email == userEmailReqBody.NewEmail {
		c.JSON(http.StatusBadRequest, gin.H{"error": "New email is the same as the current email"})
		return
	}

	updateEmail := bson.M{"$set": bson.M{"email": userEmailReqBody.NewEmail}}

	_, err = collection.UpdateOne(context.Background(), FindUser, updateEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update email"})
		return
	}

	c.Next()
}









