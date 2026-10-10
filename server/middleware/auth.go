package middleware

import (
    "net/http"
    "strconv"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

func Auth(secret string) gin.HandlerFunc {
    return func(c *gin.Context) {
        parts := strings.Fields(c.GetHeader("Authorization"))

        if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error": "Missing bearer token",
            })
            return
        }

        claims := &jwt.RegisteredClaims{}

        token, err := jwt.ParseWithClaims(parts[1], claims,
            func(token *jwt.Token) (any, error) {
                if token.Method != jwt.SigningMethodHS256 {
                    return nil, jwt.ErrSignatureInvalid
                }
                return []byte(secret), nil
            },
        )

        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error": "Invalid or expired token",
            })
            return
        }

        userID, err := strconv.ParseUint(claims.Subject, 10, 64)
        if err != nil || userID == 0 {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "error": "Invalid token subject",
            })
            return
        }

        c.Set("userID", uint(userID))
        c.Next()
    }
}
