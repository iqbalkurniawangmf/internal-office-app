package routes

import (
    "internal-office-backend/internal/middleware"
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {

    // Health check, tidak pakai auth
    rg.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status":  "ok",
            "message": "Internal Office Management API Running",
        })
    })

    // Department routes, pakai Auth
    deptGroup := rg.Group("/departments")
    deptGroup.Use(middleware.AuthMiddleware()) 
    RegisterDepartmentRoutes(deptGroup, db)
}