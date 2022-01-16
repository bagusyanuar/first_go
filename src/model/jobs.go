package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Jobs struct {
	ID           uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	CompaniesID  uuid.UUID `gorm:"column:company_id;type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;not null" json:"company_id"`
	ProfessionID uuid.UUID `gorm:"column:profession_id;type:int(11);not null" json:"profession_id"`
	Title        string    `gorm:"type:varchar(255);not null;" json:"title"`
	Slug         string    `gorm:"type:varchar(255);not null;" json:"slug"`
	ExpiredAt    time.Time `gorm:"column:expired_at;not null" json:"expired_at"`
	CreatedAt    time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

func (j *Jobs) BeforCreate(tx *gorm.DB) (err error) {
	j.ID = uuid.New()
	return
}

func (Jobs) TableName() string {
	return "jobs"
}
