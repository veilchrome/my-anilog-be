package utils

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secretKey string
}

func NewJWTManager(secret string) *JWTManager {
	return &JWTManager{secretKey: secret}
}

func (j *JWTManager) GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTManager) VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userID, ok := claims["sub"].(string); ok {
			return userID, nil
		}
	}
	return "", jwt.ErrTokenInvalidClaims
}

func AuthMiddleware(jwtManager *JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		userID, err := jwtManager.VerifyToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		c.Set("userID", userID)
		c.Next()
	}
}
