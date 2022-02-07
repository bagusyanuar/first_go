package model

import (
	"time"

	"gorm.io/gorm"
)


type Grade struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Slug      string    `gorm:"type:varchar(255);not null" json:"slug"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

func (grade *Grade) BeforeCreate(tx *gorm.DB) (err error) {
	grade.CreatedAt = time.Now()
	grade.UpdatedAt = time.Now()
	return
}

func (Grade) TableName() string {
	return "grades"
}

type GradeSimple struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Slug      string    `gorm:"type:varchar(255);not null" json:"slug"`
}

func (GradeSimple) TableName() string {
	return "grades"
}