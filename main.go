package main

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

// ClientIPMiddleware extracts the client's real IP address
func ClientIPMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        clientIP := c.ClientIP()
        if forwarded := c.GetHeader("X-Forwarded-For"); forwarded != "" {
            // X-Forwarded-For header may contain multiple IP addresses, the first one is the client's real IP
            clientIP = strings.Split(forwarded, ",")[0]
        }
        // Set the client's real IP address to the context
        c.Set("ClientIP", clientIP)
        c.Next()
    }
}

func main() {
    r := gin.Default()

    // Apply the middleware
    r.Use(ClientIPMiddleware())

    r.GET("/", func(c *gin.Context) {
        clientIP, _ := c.Get("ClientIP")
        c.JSON(http.StatusOK, gin.H{"ClientIP": clientIP})
    })

    r.Run(":5000")
}
