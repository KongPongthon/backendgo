package handlers

import (
	"fmt"
	"net/http"

	"backendgo/db"
	"backendgo/models"
	"backendgo/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register user
func RegisterHandler(c *gin.Context) {
    fmt.Println("RegisterHandler",c.Request.Body)
    var body struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    fmt.Printf("Request body: %+v\n", body)
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    // hash password
    hashed, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
    user := models.User{Username: body.Username, Password: string(hashed)}

    if err := db.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

// Login user
func LoginHandler(c *gin.Context) {
    var body struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    var user models.User
    if err := db.DB.First(&user, "username = ?", body.Username).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    token, _ := utils.GenerateToken(user.Username)

    // ✅ ส่ง JWT กลับไปใน JSON ได้ (ใช้ axios ได้เลย)
    c.JSON(http.StatusOK, gin.H{"token": token})

    // ✅ หรือถ้าอยากเก็บใน Cookie (Secure/HttpOnly)
    // c.SetCookie("auth_token", token, 3600*24, "/", "localhost", false, true)
}
