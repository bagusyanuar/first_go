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

func FindSubjectBySlug(subject *model.Subject, slug string) (data *model.Subject, err error) {
	if err = database.DATABASE.Debug().Where("slug = ?", slug).First(&subject).Error; err != nil {
		return nil, err
	}

	return subject, nil
}
