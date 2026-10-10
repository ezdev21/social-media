package controllers

import (
    "net/http"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"

    "server/models"
	"strconv"
)

type AuthController struct {
    DB        *gorm.DB
    JWTSecret []byte
}

type RegisterInput struct {
    Name     string `json:"name" binding:"required,min=2,max=100"`
    Email    string `json:"email" binding:"required,email,max=255"`
    Password string `json:"password" binding:"required,min=8"`
}

type LoginInput struct {
    Email    string `json:"email" binding:"required,email,max=255"`
    Password string `json:"password" binding:"required"`
}

func NewAuthController(db *gorm.DB, secret string) *AuthController {
    return &AuthController{
        DB:        db,
        JWTSecret: []byte(secret),
    }
}

func (a *AuthController) Register(c *gin.Context) {
    var input RegisterInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    input.Name = strings.TrimSpace(input.Name)
    input.Email = strings.ToLower(strings.TrimSpace(input.Email))

    if input.Name == "" || len(input.Password) > 72 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid name or password"})
        return
    }

    hash, err := bcrypt.GenerateFromPassword(
        []byte(input.Password), bcrypt.DefaultCost,
    )
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
        return
    }

    user := models.User{
        Name:         input.Name,
        Email:        input.Email,
        PasswordHash: string(hash),
    }

    if err := a.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusConflict, gin.H{"error": "Could not register this email"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (a *AuthController) Login(c *gin.Context) {
    var input LoginInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    input.Email = strings.ToLower(strings.TrimSpace(input.Email))

    var user models.User
    if err := a.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    if bcrypt.CompareHashAndPassword(
        []byte(user.PasswordHash), []byte(input.Password),
    ) != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    claims := jwt.RegisteredClaims{
        Subject:   strconv.FormatUint(uint64(user.ID), 10),
        IssuedAt:  jwt.NewNumericDate(time.Now()),
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString(a.JWTSecret)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "access_token": signedToken,
        "token_type":   "Bearer",
        "expires_in":   900,
        "user":         user,
    })
}