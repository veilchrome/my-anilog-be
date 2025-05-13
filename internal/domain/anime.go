package domain

type Anime struct {
	ID       uint   `gorm:"primaryKey"`
	UserID   string `gorm:"type:char(36);index"`
	MalID    int
	Title    string
	Status   string // favorite, watching, watched
	ImageURL string
}
