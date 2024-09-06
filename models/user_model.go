package models

type User struct {
	ID    uint   `gorm:"primarykey" json:"id"`
	Name  string `gorm:"not null;unique" json:"name"`
	Email string `gorm:"not null;unique" json:"email"`
}
