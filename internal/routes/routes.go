package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/veilchrome/myanilog-be/internal/handler"
	"github.com/veilchrome/myanilog-be/internal/service"
	"github.com/veilchrome/myanilog-be/internal/utils"
)

func RegisterRoutes(r *gin.Engine, userService service.UserService, animeService service.AnimeService, jwt *utils.JWTManager) {
	// Auth Routes
	handler.RegisterAuthRoutes(r, userService, jwt)

	// Anime Routes
	handler.RegisterAnimeRoutes(r, animeService, jwt)

	// Protected user info
	r.GET("/me", utils.AuthMiddleware(jwt), handler.MeHandler)
}
