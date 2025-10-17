package handlers

import (
	"backendgo/db"
	"backendgo/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProfileHandler(c *gin.Context) {
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    //! เตรียมฐานข้อมูล
    var userDetail models.UserDetail
    
    // ดึงข้อมูล
    if err := db.DB.First(&userDetail, "id = ?", userID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    fmt.Println("userID", userDetail)
    c.JSON(http.StatusOK, gin.H{
        "message": "Welcome to profile!",
        "userID":  userID,
        "Code":    userDetail.Code,
        "FirstName":   userDetail.FirstName,
        "LastName":    userDetail.LastName,
        "Email":       userDetail.Email,
        "Phone":       userDetail.Phone,
        "Position":    userDetail.Position,
        "Department":  userDetail.Department,
        "Status":      userDetail.Status,
    })
}
