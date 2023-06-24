package admin

import (
	"codename_backend/database"
	"codename_backend/models"
	"context"

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
		c.JSON(500, gin.H{
			"message": "Something went wrong",
		})
		return
	}
	defer cursor.Close(context.Background())

	var users []models.User
	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			c.JSON(500, gin.H{
				"message": "Something went wrong",
			})
			return
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(500, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(200, users)
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
}