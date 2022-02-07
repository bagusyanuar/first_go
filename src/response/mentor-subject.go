package response

import "first_go/src/model"

type PreloadMentorSubjectAll struct {
	model.MentorSubject
	Subject *Subject `gorm:"foreignKey:SubjectID" json:"subject"`
	Grade   *Grade   `gorm:"foreignKey:GradeID" json:"grade"`
}

type Subject struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Slug      string    `gorm:"type:varchar(255);not null" json:"slug"`
	Icon      *string   `gorm:"type:text;" json:"icon"`
}

type Grade struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Slug      string    `gorm:"type:varchar(255);not null" json:"slug"`
}
