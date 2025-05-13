package repository

import (
	"github.com/veilchrome/myanilog-be/internal/domain"
	"gorm.io/gorm"
)

type AnimeRepository interface {
	Create(anime *domain.Anime) error
}

type animeRepository struct {
	db *gorm.DB
}

func NewAnimeRepository(db *gorm.DB) AnimeRepository {
	return &animeRepository{db: db}
}

func (r *animeRepository) Create(anime *domain.Anime) error {
	return r.db.Create(anime).Error
}
