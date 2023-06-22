package main

import (
	"codename_backend/database"
	"codename_backend/socketio"
	"io"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := os.MkdirAll("Logs", os.ModePerm)
	if err != nil {
		log.Fatal("Failed to create Logs directory: ", err)
	}

	logFilePath := filepath.Join("Logs", "program.log")
	logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Failed to create log file: ", err)
	}
	defer logFile.Close()

	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	log.SetFormatter(&log.JSONFormatter{})

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {

		log.WithFields(log.Fields{
			"route": "/",
		}).Info("Received GET request")

		c.JSON(200, gin.H{
			"message": "welcome to codename backend",
		})
	})
	router.GET("/socket", gin.WrapH(socketio.SocketIO()))
	
	database.ConnectMongoDB()

	router.Run(":8080")
}
