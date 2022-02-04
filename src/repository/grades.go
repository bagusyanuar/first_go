package repository

import (
	"first_go/database"
	"first_go/src/model"
)


func FindAllGrades(model *[]model.Grade, param string) (data *[]model.Grade, err error) {
	if err = database.DATABASE.Debug().Where("name LIKE ?", "%"+param+"%").Find(&model).Error; err != nil {
		return model, err
	}
	return model, nil
}