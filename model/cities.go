package model

import "time"

type Cities struct {
	ID        uint      `gorm:"type:int(11);primaryKey;" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Slug      string    `gorm:"type:varchar(255);not null" json:"slug"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

func (Cities) TableName() string {
	return "cities"
}
