package middleware

import (
	"backendgo/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func OriginCheckMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        origin := c.GetHeader("Origin")
        allowed := false
        // for _, o := range config.AllowedOrigins {
        //     if strings.HasPrefix(origin, o) {
        //         allowed = true
        //         break
        //     }
        // }
        if origin == "" { // ถ้าไม่มี origin เช่น Postman
            allowed = true
        } else {
            for _, o := range config.AllowedOrigins {
                if strings.HasPrefix(origin, o) {
                    allowed = true
                    break
                }
            }
        }

        if !allowed {
            c.JSON(http.StatusForbidden, gin.H{"error": "Origin not allowed"})
            c.Abort()
            return
        }

        c.Next()
    }
}
