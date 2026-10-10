package routes

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "server/controllers"
    authMiddleware "server/middleware"
	"server/models"
)

func Setup(db *gorm.DB, secret string) *gin.Engine {
    r := gin.Default()

    auth := controllers.NewAuthController(db, secret)

    authRoutes := r.Group("/auth")
    {
        authRoutes.POST("/register", auth.Register)
        authRoutes.POST("/login", auth.Login)
    }

    protected := r.Group("/")
    protected.Use(authMiddleware.Auth(secret))
    {
        protected.GET("/me", func(c *gin.Context) {
            userID := c.GetUint("userID")

            var user models.User
            if err := db.First(&user, userID).Error; err != nil {
                c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
                return
            }

            c.JSON(http.StatusOK, gin.H{"user": user})
        })
    }

    return r
}