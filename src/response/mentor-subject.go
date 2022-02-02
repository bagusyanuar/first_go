package response

import "first_go/src/model"

type PreloadMentorSubjectWSubject struct {
	model.MentorSubject
	Subject *model.Subject `gorm:"foreignKey:SubjectID" json:"subject"`
}
