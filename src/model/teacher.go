package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Teacher struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey;" json:"id"`
	UserID    uuid.UUID `gorm:"column:user_id;type:char(36);not null" json:"user_id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Phone     string    `gorm:"type:varchar(16);not null;" json:"phone"`
	Avatar    *string   `gorm:"type:text;" json:"avatar"`
	Address   *string   `gorm:"type:text;" json:"address"`
	Gender    uint8     `gorm:"type:smallint;not null;default:0" json:"gender"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

type MigrationTeacher struct {
	Teacher
	User User `gorm:"foreignKey:UserID"`
}
func (t *Teacher) BeforCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}

func (Teacher) TableName() string {
	return "teachers"
}
