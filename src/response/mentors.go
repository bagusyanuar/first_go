package response

import (
	"first_go/src/model"
)

// used on endpoint [
// 	"/mentor/me",
// ]

type MentorProfileResponse struct {
	model.Mentor
	User   PreloadUser                     `gorm:"foreignKey:UserID" json:"account"`
	Skills []*PreloadMentorSubjectWSubject `gorm:"foreignKey:MentorID" json:"skills"`
}
