package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Member struct {
	ID          uuid.UUID  `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	UserID      uuid.UUID  `gorm:"column:user_id;type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;not null" json:"user_id"`
	Name        string     `gorm:"type:varchar(255);not null;" json:"name"`
	Phone       string     `gorm:"type:varchar(16);not null;" json:"phone"`
	Avatar      string     `gorm:"type:text;not null;" json:"avatar"`
	Address     string     `gorm:"type:text;not null;" json:"address"`
	Gender      uint8      `gorm:"type:smallint;not null;comment:0 Not Set 1 Laki-Laki 2 Perempuan;default:0" json:"gender"`
	DateOfBirth *time.Time `gorm:"type:date;" json:"date_of_birth"`
	CreatedAt   time.Time  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;not null" json:"updated_at"`
}

type MigrationMember struct {
	Member
	User User `gorm:"foreignKey:UserID"`
}

func (member *Member) BeforeCreate(tx *gorm.DB) (err error) {
	member.ID = uuid.New()
	member.CreatedAt = time.Now()
	member.UpdatedAt = time.Now()
	return
}

func (Member) TableName() string {
	return "members"
}
