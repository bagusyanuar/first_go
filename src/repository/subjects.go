package repository

import (
	"first_go/database"
	"first_go/src/model"
)

func FindSubjects(subject *[]model.Subject, param string) (s *[]model.Subject, err error) {
	if err = database.DATABASE.Debug().Where("name LIKE ?", "%"+param+"%").Find(&subject).Error; err != nil {
		return subject, err
	}
	return subject, nil
}
