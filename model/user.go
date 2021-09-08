package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Email     string         `gorm:"type:varchar(255);not null" json:"email"`
	Password  *string        `gorm:"type:text" json:"password"`
	Roles     datatypes.JSON `gorm:"type:longtext;not null" json:"roles"`
	Provider  datatypes.JSON `gorm:"type:longtext;not null" json:"provider"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
}

func (u *User) BeforCreate(tx *gorm.DB) (err error)  {
	u.ID = uuid.New()
	return
}

func (User) TableName() string  {
	return "users"
}
