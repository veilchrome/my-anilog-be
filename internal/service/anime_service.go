package service

import (
	"github.com/veilchrome/myanilog-be/internal/domain"
	"github.com/veilchrome/myanilog-be/internal/repository"
)

type AnimeService interface {
	SaveFavorite(userID string, malID int, title, status, imageURL string) error
	GetUserAnimeList(userID string) ([]domain.UserAnimeList, error)
	UpdateUserAnime(userID string, malID int, status, note string) error
	DeleteUserAnime(userID string, malID int) error
}

type animeService struct {
	repo         repository.AnimeRepository
	userListRepo repository.UserAnimeListRepository
}

func NewAnimeService(repo repository.AnimeRepository, userListRepo repository.UserAnimeListRepository) AnimeService {
	return &animeService{
		repo:         repo,
		userListRepo: userListRepo,
	}
}

func (s *animeService) GetUserAnimeList(userID string) ([]domain.UserAnimeList, error) {
	return s.userListRepo.GetByUserID(userID)
}

func (s *animeService) UpdateUserAnime(userID string, malID int, status, note string) error {
	return s.userListRepo.UpdateStatus(userID, malID, status, note)
}

func (s *animeService) DeleteUserAnime(userID string, malID int) error {
	return s.userListRepo.Delete(userID, malID)
}

func (s *animeService) SaveFavorite(userID string, malID int, title, status, imageURL string) error {
	anime := &domain.Anime{
		UserID:   userID,
		MalID:    malID,
		Title:    title,
		Status:   status,
		ImageURL: imageURL,
	}
	return s.repo.Create(anime)
}
