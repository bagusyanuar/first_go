package response

import "github.com/google/uuid"

type PreloadUser struct {
	ID       uuid.UUID
	Email    string
	Username string
}

func (PreloadUser) TableName() string {
	return "users"
}
