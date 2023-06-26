package admin

import (
	"codename_backend/database"
	"codename_backend/models"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	DB_NAME        = "codename"
	CollectionName = "users"
)

func GetALLUSERS(c *gin.Context) {
	collection := database.GetCollection(CollectionName)
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}
	defer cursor.Close(context.Background())

	var users []models.User
	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Something went wrong",
			})
			return
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	if len(users) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No users found",
		})
		return
	}

	currentTimestamp := time.Now().Format("2006-01-02_15-04-05")
	fileName := "data/" + "users_" + currentTimestamp + ".json"
	data, err := json.Marshal(users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to marshal data",
		})
		return
	}

	err = ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to write data to file",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

// func GetASingleUserFromID(c *gin.Context) {
// 	collection := database.GetCollection(CollectionName)
// 	var user models.User
// 	type IDReq struct {
// 		ID string `json:"id"`
// 	}
// 	var idReq IDReq
// 	err := c.BindJSON(&idReq)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Invalid request",
// 		})
// 		return
// 	}
// 	err = collection.FindOne(context.Background(), bson.M{"id": idReq.ID}).Decode(&user)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Invalid request",
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, user)
// 	c.Next()
// }

// count new users todayUser - yesterdayUser

func NewUsers(c *gin.Context) {
	collection := database.GetCollection(CollectionName)
	var usersYesterday models.User
	var usersToday models.User

	yesterdayUser := collection.FindOne(context.Background(), bson.M{"created_at": time.Now().AddDate(0, 0, -1).Format("2006-01-02")})
	todayUser := collection.FindOne(context.Background(), bson.M{"created_at": time.Now().Format("2006-01-02")})

	if err := yesterdayUser.Decode(&usersYesterday); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	if err := todayUser.Decode(&usersToday); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	// newUser := len(usersToday) - len(usersYesterday)

	c.JSON(http.StatusOK, gin.H{
		"yesterdayUser": usersYesterday,
		"todayUser":     usersToday,
		// "newUser":       newUser,
	})
}
