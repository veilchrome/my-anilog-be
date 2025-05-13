package domain

type User struct {
	ID       string `gorm:"type:char(36);primaryKey"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}
