package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	UserID    uuid.UUID `gorm:"column:user_id;type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;not null" json:"user_id"`
	Name      string    `gorm:"type:varchar(255);not null;" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

type MigrationAdmin struct {
	Member
	User User `gorm:"foreignKey:UserID"`
}

func (a *Admin) BeforCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}

func (Admin) TableName() string {
	return "admins"
}
