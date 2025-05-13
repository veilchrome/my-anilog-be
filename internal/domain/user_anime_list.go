package domain

import "time"

type UserAnimeList struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	UserID    string `gorm:"type:char(36);not null"`
	MalID     int    `gorm:"not null"`
	Status    string `gorm:"type:enum('watched','watching','favorite');not null"`
	Note      string
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
