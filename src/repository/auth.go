package repository

import (
	"first_go/database"
	"first_go/src/lib"
	"first_go/src/model"
)

func SignInMember(user *model.UserMember, email string, password string, provider string) (u *model.UserMember, err error) {
	if err = database.DATABASE.Debug().Preload("Member").Joins("JOIN members ON users.id = members.user_id").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	//if member using app provider do check password
	if provider == "app" && !lib.IsPasswordValid(password, *user.Password) {
		return user, lib.ErrorInvalidPassword
	}

	return user, nil
}

func SignInMentor(user *model.UserMentor, email string, password string, provider string) (u *model.UserMentor, err error) {
	if err = database.DATABASE.Debug().Preload("Mentor").Joins("JOIN mentors ON users.id = mentors.user_id").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	//if mentor using app provider do check password
	if provider == "app" && !lib.IsPasswordValid(password, *user.Password) {
		return user, lib.ErrorInvalidPassword
	}

	return user, nil
}
