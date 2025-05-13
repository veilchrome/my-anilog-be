package service

import (
	"github.com/veilchrome/myanilog-be/internal/domain"
	"github.com/veilchrome/myanilog-be/internal/repository"
)

type UserAnimeListService interface {
	AddAnime(userID string, malID int, status string, note string) error
	ListAnime(userID string) ([]domain.UserAnimeList, error)
	UpdateAnime(userID string, malID int, status string, note string) error
	DeleteAnime(userID string, malID int) error
}

type userAnimeListService struct {
	repo repository.UserAnimeListRepository
}

func NewUserAnimeListService(repo repository.UserAnimeListRepository) UserAnimeListService {
	return &userAnimeListService{repo: repo}
}

func (s *userAnimeListService) AddAnime(userID string, malID int, status string, note string) error {
	item := &domain.UserAnimeList{
		UserID: userID,
		MalID:  malID,
		Status: status,
		Note:   note,
	}
	return s.repo.Add(item)
}

func (s *userAnimeListService) ListAnime(userID string) ([]domain.UserAnimeList, error) {
	return s.repo.GetByUserID(userID)
}

func (s *userAnimeListService) UpdateAnime(userID string, malID int, status string, note string) error {
	return s.repo.UpdateStatus(userID, malID, status, note)
}

func (s *userAnimeListService) DeleteAnime(userID string, malID int) error {
	return s.repo.Delete(userID, malID)
}
