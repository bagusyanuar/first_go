package model

import (
	"time"

	"gorm.io/gorm"
)

type Subject struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Slug      string    `gorm:"type:varchar(255);not null" json:"slug"`
	Icon      *string   `gorm:"type:text;" json:"icon"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

func (subject *Subject) BeforeCreate(tx *gorm.DB) (err error) {
	subject.CreatedAt = time.Now()
	subject.UpdatedAt = time.Now()
	return
}

func (Subject) TableName() string {
	return "subjects"
}

type SubjectSimple struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Slug      string    `gorm:"type:varchar(255);not null" json:"slug"`
	Icon      *string   `gorm:"type:text;" json:"icon"`
}

func (SubjectSimple) TableName() string {
	return "subjects"
}
