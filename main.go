package main

import (
	"codename_backend/database"
	"codename_backend/middlewares"
	"codename_backend/socketio"
	"codename_backend/utils"
	"io"
	"net/http"
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

	// Middleware function for logging
	router.Use(func(c *gin.Context) {
		log.WithFields(log.Fields{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
		}).Info("Received request")
		c.Next()
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome to codename backend",
		})
	})

	router.GET("/socket.io/*any", gin.WrapH(socketio.SocketIO()))

	// router.GET("/codename", func(c *gin.Context) {
	// 	log.WithFields(log.Fields{
	// 		"route": "/codename",
	// 	}).Info("Received GET request")
	// 	RandomCodeName := utils.GetCodeName()
	// 	c.JSON(200, gin.H{
	// 		"codename": RandomCodeName,
	// 	})
	// })
	router.GET("/codename", middlewares.CheckLoggedIn, func(ctx *gin.Context) {
		log.WithFields(log.Fields{
			"route": "/codename",
		}).Info(
			"Received GET request",
		)
		RandomCodeName := utils.GetCodeName()
		ctx.JSON(200, gin.H{
			"codename": RandomCodeName,
		})

	})

	router.POST("/register", middlewares.RegisterNewUsers, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "User registered successfully",
		})
	})

	router.POST("/login", middlewares.LoginMiddleware, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "User logged in successfully",
		})
	})

	router.POST("/forget" , middlewares.ResetPassword, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OTP sent successfully",
		})
	})

	router.POST("/reset", middlewares.VerifyPassword, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Password reset successfully",
		})
	})

	database.ConnectMongoDB()

	router.Run(":8080")
}
