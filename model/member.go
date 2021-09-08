package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Member struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey;" json:"id"`
	UserID    uuid.UUID `gorm:"column:user_id;type:char(36);not null" json:"user_id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Phone     string    `gorm:"type:varchar(16);not null;" json:"phone"`
	Avatar    *string   `gorm:"type:text;" json:"avatar"`
	Address   *string   `gorm:"type:text;" json:"address"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

type MigrationMember struct {
	Member
	User User `gorm:"foreignKey:UserID"`
}

func (m *Member) BeforCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	return
}

func (Member) TableName() string {
	return "members"
}
