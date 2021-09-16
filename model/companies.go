package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Companies struct {
	ID           uuid.UUID `gorm:"type:char(36);primaryKey;" json:"id"`
	UserID       uuid.UUID `gorm:"column:user_id;type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;not null" json:"user_id"`
	Name         string    `gorm:"type:varchar(255);not null" json:"name"`
	Slug         string    `gorm:"type:varchar(255);not null" json:"slug"`
	IndustrialID uint      `gorm:"column:industrial_id;type:int(11);not null" json:"industrial_id"`
	Phone        string    `gorm:"type:varchar(16);not null;" json:"phone"`
	Avatar       string    `gorm:"type:text;not null;" json:"avatar"`
	Address      string    `gorm:"type:text;not null;" json:"address"`
	CreatedAt    time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

func (c *Companies) BeforCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}

func (Companies) TableName() string {
	return "companies"
}
