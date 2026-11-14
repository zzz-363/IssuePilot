package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/api/v1/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "code":    200,
            "message": "ok",
            "data": gin.H{
                "service": "issuepilot",
                "status":  "healthy",
            },
        })
    })

    if err := r.Run(":8080"); err != nil {
        panic(err)
    }
}