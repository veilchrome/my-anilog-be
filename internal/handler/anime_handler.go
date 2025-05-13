package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/veilchrome/myanilog-be/internal/service"
	"github.com/veilchrome/myanilog-be/internal/utils"
)

type FavoriteRequest struct {
	MalID    int    `json:"mal_id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Status   string `json:"status" binding:"required"` // favorite / watching / watched
	ImageURL string `json:"image_url"`
}

type UpdateRequest struct {
	MalID  int    `json:"mal_id" binding:"required"`
	Status string `json:"status" binding:"required"`
	Note   string `json:"note"`
}

func RegisterAnimeRoutes(r *gin.Engine, animeService service.AnimeService, jwt *utils.JWTManager) {
	anime := r.Group("/anime")
	anime.Use(utils.AuthMiddleware(jwt))
	{
		anime.GET("/search", searchAnimeHandler)

		anime.POST("/favorite", func(c *gin.Context) {
			userID := c.MustGet("userID").(string)

			var req FavoriteRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if err := animeService.SaveFavorite(userID, req.MalID, req.Title, req.Status, req.ImageURL); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Anime saved successfully"})
		})

		anime.GET("/list", func(c *gin.Context) {
			userID := c.MustGet("userID").(string)

			list, err := animeService.GetUserAnimeList(userID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch anime list"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"anime_list": list})
		})

		anime.PUT("/", func(c *gin.Context) {
			userID := c.MustGet("userID").(string)

			var req UpdateRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			err := animeService.UpdateUserAnime(userID, req.MalID, req.Status, req.Note)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Anime updated successfully"})
		})

		anime.DELETE("/", func(c *gin.Context) {
			userID := c.MustGet("userID").(string)

			malIDStr := c.Query("mal_id")
			if malIDStr == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "mal_id query parameter required"})
				return
			}

			malID, err := strconv.Atoi(malIDStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "mal_id must be a valid integer"})
				return
			}

			err = animeService.DeleteUserAnime(userID, malID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Anime deleted successfully"})
		})
	}
}

func searchAnimeHandler(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing query parameter 'q'"})
		return
	}

	url := fmt.Sprintf("https://api.jikan.moe/v4/anime?q=%s", query)
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch anime from API"})
		return
	}
	defer resp.Body.Close()

	var result interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response"})
		return
	}

	c.JSON(http.StatusOK, result)
}
