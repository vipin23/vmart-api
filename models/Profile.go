package models

type Profile struct {
	BaseModel
	Title     string `gorm:"size:255;not null;unique" json:"title"`
	FirstName string `gorm:"size:255;not null;" json:"first_name"`
	LastName  string `gorm:"size:255;not null;" json:"last_name"`
	UserID    uint64 `gorm:"not null" json:"user_id"`
	User      User
}
