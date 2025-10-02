package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func ProfileHandler(c *gin.Context) {
    userID, _ := c.Get("userID")
    c.JSON(http.StatusOK, gin.H{
        "message": "Welcome to profile!",
        "userID":  userID,
    })
}
