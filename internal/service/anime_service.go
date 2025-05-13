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

// type animeService struct {
// 	repo         repository.AnimeRepository
// 	userListRepo repository.UserAnimeListRepository
// }

// func NewAnimeService(repo repository.AnimeRepository, userListRepo repository.UserAnimeListRepository) AnimeService {
// 	return &animeService{
// 		repo:         repo,
// 		userListRepo: userListRepo,
// 	}
// }

type animeService struct {
	repo          repository.AnimeRepository
	userAnimeRepo repository.UserAnimeListRepository
}

func NewAnimeService(repo repository.AnimeRepository, userAnimeRepo repository.UserAnimeListRepository) AnimeService {
	return &animeService{repo: repo, userAnimeRepo: userAnimeRepo}
}

// func (s *animeService) GetUserAnimeList(userID string) ([]domain.UserAnimeList, error) {
// 	return s.userListRepo.GetByUserID(userID)
// }

func (s *animeService) UpdateUserAnime(userID string, malID int, status, note string) error {
	return s.userAnimeRepo.UpdateStatus(userID, malID, status, note)
}

func (s *animeService) DeleteUserAnime(userID string, malID int) error {
	return s.userAnimeRepo.Delete(userID, malID)
}

// func (s *animeService) SaveFavorite(userID string, malID int, title, status, imageURL string) error {
// 	anime := &domain.Anime{
// 		UserID:   userID,
// 		MalID:    malID,
// 		Title:    title,
// 		Status:   status,
// 		ImageURL: imageURL,
// 	}
// 	return s.repo.Create(anime)
// }

func (s *animeService) SaveFavorite(userID string, malID int, title, status, imageURL string) error {
	anime := &domain.Anime{
		UserID:   userID,
		MalID:    malID,
		Title:    title,
		Status:   status,
		ImageURL: imageURL,
	}
	if err := s.repo.Create(anime); err != nil {
		return err
	}

	// Save also to user_anime_lists
	return s.userAnimeRepo.Add(&domain.UserAnimeList{
		UserID: userID,
		MalID:  malID,
		Status: status,
		Note:   "",
	})
}

func (s *animeService) GetUserAnimeList(userID string) ([]domain.UserAnimeList, error) {
	return s.userAnimeRepo.GetByUserID(userID)
}
