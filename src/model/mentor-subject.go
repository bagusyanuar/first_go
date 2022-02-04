package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MentorSubject struct {
	ID        uuid.UUID `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	MentorID  uuid.UUID `gorm:"column:mentor_id;type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;not null" json:"mentor_id"`
	SubjectID uint      `gorm:"column:subject_id;type:bigint(20) UNSIGNED;not null" json:"subject_id"`
	GradeID   uint      `gorm:"column:grade_id;type:bigint(20) UNSIGNED;not null" json:"grade_id"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

func (ms *MentorSubject) BeforeCreate(tx *gorm.DB) (err error) {
	ms.ID = uuid.New()
	ms.CreatedAt = time.Now()
	ms.UpdatedAt = time.Now()
	return
}

func (MentorSubject) TableName() string {
	return "mentor_subject"
}
