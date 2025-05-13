package handler

import (
	"net/http"

	"github.com/veilchrome/myanilog-be/internal/service"
	"github.com/veilchrome/myanilog-be/internal/utils"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

func RegisterAuthRoutes(r *gin.Engine, userService service.UserService, jwtManager *utils.JWTManager) {
	r.POST("/register", func(c *gin.Context) {
		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := userService.Register(req.Username, req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"user": user})
	})

	r.POST("/login", func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := userService.LoginWithIdentifier(req.Identifier, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		token, err := jwtManager.GenerateToken(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"user": gin.H{
				"id":       user.ID,
				"username": user.Username,
				"email":    user.Email,
			},
		})
	})
}

func MeHandler(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userID not found in context"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user_id": userID})
}
