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

	currentTimestamp := time.Now().Format("2006-01-02_15-04-05")
	fileName := "users_" + currentTimestamp + ".json"
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

func GetASingleUserFromID(c *gin.Context) {
	collection := database.GetCollection(CollectionName)
	var user models.User
	if err := collection.FindOne(context.Background(), bson.M{"_id": c.Param("id")}).Decode(&user); err != nil {
		c.JSON(500, gin.H{
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(200, user)

	c.Next()
}