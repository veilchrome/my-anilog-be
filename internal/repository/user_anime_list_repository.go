package repository

import (
	"github.com/veilchrome/myanilog-be/internal/domain"
	"gorm.io/gorm"
)

type UserAnimeListRepository interface {
	Add(item *domain.UserAnimeList) error
	GetByUserID(userID string) ([]domain.UserAnimeList, error)
	UpdateStatus(userID string, malID int, status string, note string) error
	Delete(userID string, malID int) error
}

type userAnimeListRepository struct {
	db *gorm.DB
}

func NewUserAnimeListRepository(db *gorm.DB) UserAnimeListRepository {
	return &userAnimeListRepository{db: db}
}

func (r *userAnimeListRepository) Add(item *domain.UserAnimeList) error {
	return r.db.Create(item).Error
}

func (r *userAnimeListRepository) GetByUserID(userID string) ([]domain.UserAnimeList, error) {
	var list []domain.UserAnimeList
	err := r.db.Where("user_id = ?", userID).Find(&list).Error
	return list, err
}

func (r *userAnimeListRepository) UpdateStatus(userID string, malID int, status string, note string) error {
	return r.db.Model(&domain.UserAnimeList{}).
		Where("user_id = ? AND mal_id = ?", userID, malID).
		Updates(map[string]interface{}{"status": status, "note": note}).Error
}

func (r *userAnimeListRepository) Delete(userID string, malID int) error {
	return r.db.Where("user_id = ? AND mal_id = ?", userID, malID).Delete(&domain.UserAnimeList{}).Error
}
