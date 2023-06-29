package main

import (
	"codename_backend/admin"
	"codename_backend/auth"
	"codename_backend/database"
	"codename_backend/middlewares"

	// "codename_backend/routes"
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
	authOk := auth.BasicAuth()
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
	router.GET("/codename", func(ctx *gin.Context) {
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

	router.POST("/forget", authOk, middlewares.ResetPassword, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OTP sent successfully",
		})
	})

	router.POST("/reset", middlewares.VerifyPassword, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Password reset successfully",
		})
	})

	router.POST("/newpwd", middlewares.ChangePassword, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Password changed successfully",
		})
	})
	router.GET("/admin", auth.JWT_BASIC_AUTH(), admin.GetALLUSERS, func(c *gin.Context) {
		admin.GetALLUSERS(c)
	})

	// router.POST("/id" , admin.GetASingleUserFromID , func(c *gin.Context) {
	// 	admin.GetASingleUserFromID(c)
	// })

	router.POST("/changeusername", middlewares.UserNameChange, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Username changed successfully",
		})
	})
	router.GET("/newuser", admin.NewUsers, func(c *gin.Context) {
		admin.NewUsers(c)
	})

	router.POST("/deleteuser", middlewares.UserDelete, func(c *gin.Context) {
		middlewares.UserDelete(c)
	})
	database.ConnectMongoDB()
	database.ConnectRedisDB()
	// JWT Generation
	TOKEN, err := auth.GenerateJWT("admin")
	if err != nil {
		log.Fatal("Failed to generate JWT: ", err)
	}
	log.Info("JWT: ", TOKEN)
	// save the token in a file for later use
	err = os.WriteFile("token/token.txt", []byte("Bearer" + " " + TOKEN), 0644)
	if err != nil {
		log.Fatal("Failed to save JWT: ", err)
	}

	router.Run(":8080")
}
