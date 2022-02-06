package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Email     string         `gorm:"index:idx_email,unique;type:varchar(255);not null" json:"email"`
	Username  string         `gorm:"index:idx_username,unique;type:varchar(255);not null" json:"username"`
	Password  *string        `gorm:"type:text" json:"password"`
	Roles     datatypes.JSON `gorm:"type:longtext;not null" json:"roles"`
	Provider  datatypes.JSON `gorm:"type:longtext;not null" json:"provider"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	return
}

func (User) TableName() string {
	return "users"
}

//associate with admin profile
type UserAdmin struct {
	User
	Admin Admin `gorm:"foreignKey:UserID" json:"admin"`
}

//associate with member profile
type UserMember struct {
	User
	Member Member `gorm:"foreignkey:UserID" json:"member"`
}

//associate with mentor profile
type UserMentor struct {
	User
	Mentor Mentor `gorm:"foreignkey:UserID" json:"mentor"`
}
