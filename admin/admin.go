package admin

import (
	"context"
	"codename_backend/database"
	"codename_backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetALLUSERS(c *gin.Context) {
	collection := database.GetCollection("codename") // Replace "codename" with your collection name
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
