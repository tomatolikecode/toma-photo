package global

import "time"

type Model struct {
	ID        uint `json:"id" gorm:"primarykey" `
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `json:"-" gorm:"index"`
}
