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
	Skills []*PreloadMentorSubjectAll `gorm:"foreignKey:MentorID" json:"skills"`
}

type MentorSubjectResponse struct {
	model.Mentor
	Subjects []*PreloadMentorSubjectAll `gorm:"foreignKey:MentorID" json:"subjects"`
}
