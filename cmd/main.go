// File: cmd/main.go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/veilchrome/myanilog-be/internal/repository"
	"github.com/veilchrome/myanilog-be/internal/routes"
	"github.com/veilchrome/myanilog-be/internal/service"
	"github.com/veilchrome/myanilog-be/internal/utils"
)

func main() {
	_ = godotenv.Load("../.env")

	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database")
		fmt.Println("DB_DSN is:", os.Getenv("DB_DSN"))
	}

	r := gin.Default()

	jwtManager := utils.NewJWTManager(os.Getenv("JWT_SECRET"))
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	animeRepo := repository.NewAnimeRepository(db)
	userAnimeListRepo := repository.NewUserAnimeListRepository(db)
	animeService := service.NewAnimeService(animeRepo, userAnimeListRepo)

	// userAnimeListRepo := repository.NewUserAnimeListRepository(db)
	// animeService := service.NewAnimeService(animeRepo, userAnimeListRepo)

	routes.RegisterRoutes(r, userService, animeService, jwtManager)

	r.Run()
}
