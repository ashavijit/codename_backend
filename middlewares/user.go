package middlewares

import (
	"codename_backend/database"
	"codename_backend/models"
	"codename_backend/utils"
	"context"
	"fmt"
	"time"

	// "encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	DB_NAME        = "codename"
	CollectionName = "users"
)

func RegisterNewUsers(c *gin.Context) {
	collection := database.GetCollection(CollectionName)
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
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
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		c.Abort()
		return
	}

	// Insert the user into the database
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		c.Abort()
		return
	}

	c.Next()
}

func LoginMiddleware(c *gin.Context){
	collection := database.GetCollection(CollectionName)
	var user models.User
	if err:=c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		c.Abort()
		return
	}
	filter:=bson.M{"email": user.Email , "password": user.Password}
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

func CheckLoggedIn(c *gin.Context){
	collection := database.GetCollection(CollectionName)
	var user models.User
	if err:=c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		c.Abort()
		return
	}
	filter:=bson.M{"email": user.Email , "password": user.Password}
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

func ResetPassword (c *gin.Context)  {
	collection := database.GetCollection(CollectionName)
	var user models.User
	if err:=c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		c.Abort()
		return
	}
	filter:=bson.M{"email": user.Email}
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
	_,err = collection.UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"otp": otp, "otp_timestamp": user.OTPTimestamp}})
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
        c.Abort()
        return
    }
    filter := bson.M{"email": user.Email, "otp": user.OTP}
    err := collection.FindOne(context.Background(), filter).Decode(&user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        c.Abort()
        return
    }

	if user.OTP != user.EnteredOTP || time.Now().Sub(user.OTPTimestamp).Minutes() > 1 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
		c.Abort()
		return
	}
	fmt.Println(user.OTP, user.EnteredOTP)
    _, err = collection.UpdateOne(
        context.Background(),
        filter,
        bson.M{"$set": bson.M{"password": user.Password}},
    )
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
        c.Abort()
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Password reset successful"})
    c.Next()
}
