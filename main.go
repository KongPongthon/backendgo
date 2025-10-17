package main

import (
	"backendgo/db"
	"backendgo/handlers"
	"backendgo/middleware"
	"backendgo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // เชื่อม DB
    db.ConnectDB()
    // db.DB.AutoMigrate(&models.User{}, &models.UserDetail{})
    db.DB.AutoMigrate(&models.User{}, &models.UserDetail{})


    // Middleware ทั่วไป
    r.Use(middleware.RateLimitMiddleware())
    r.Use(middleware.OriginCheckMiddleware())
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Welcome to Gin Auth API"})
    })

    // Public routes
    r.POST("/register", handlers.RegisterHandler)
    r.POST("/login", handlers.LoginHandler)

    // Protected routes
    api := r.Group("/api")
    api.Use(middleware.AuthMiddleware())
    {
        api.GET("/profile", handlers.ProfileHandler)
    }

    r.Run(":8080")
}
