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
        Code     string `json:"code"`
        FirstName string `json:"first_name"`
        LastName string `json:"last_name"`
        Email string `json:"email"`
        Phone string `json:"phone"`
        Position string `json:"position"`
        Department string `json:"department"`
        // Status models.UserStatus `json:"status"`
    }
    fmt.Printf("Request body: %+v\n", body)
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    if body.Username == "" || body.Password == "" || body.Code == "" || body.FirstName == "" || body.LastName == "" || body.Email == "" || body.Phone == "" || body.Position == "" || body.Department == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
        return
    }

    // hash password
    hashed, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }
    user := models.User{
        Username: body.Username,
        Password: string(hashed),
        UserDetail: models.UserDetail{
        Code:        body.Code,
        FirstName:   body.FirstName,
        LastName:    body.LastName,
        Email:       body.Email,
        Phone:       body.Phone,
        Position:    body.Position,
        Department:  body.Department,
        Status:      models.StatusActive,
    }}

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
    // fmt.Println("body.Username", user)

    
    // return
    if err := db.DB.Preload("UserDetail").First(&user, "username = ?", body.Username).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }
    // fmt.Println("--------------------------------------")
    // fmt.Printf("user: %+v\n", user)

    token, err := utils.GenerateToken(user.UserDetail.ID.String())
    // token, err := utils.GenerateToken(strconv.Itoa(user.ID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }
    // fmt.Println("token", token)

   // ตั้ง cookie
    c.SetCookie(
        "token",       // ชื่อ cookie
        token,         // ค่า
        3600,          // อายุ cookie เป็นวินาที (1 ชั่วโมง)
        "/",           // path
        "localhost",   // domain (ถ้าเป็น production ต้องใส่ domain จริง)
        false,         // secure (true = https)
        true,          // httpOnly (true = client JS อ่านไม่ได้)
    )

    // ✅ ส่ง JWT กลับไปใน JSON ได้ (ใช้ axios ได้เลย)
    c.JSON(http.StatusOK, gin.H{ "message": "Login success"})

    // ✅ หรือถ้าอยากเก็บใน Cookie (Secure/HttpOnly)
    // c.SetCookie("auth_token", token, 3600*24, "/", "localhost", false, true)
}
