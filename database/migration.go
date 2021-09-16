package database

import "first_go/model"

type MemberRelation struct {
	model.Member
	User model.User `gorm:"foreignKey:UserID"`
}

type CompaniesRelation struct {
	model.Companies
	User       model.User       `gorm:"foreignKey:UserID"`
	Industrial model.Industrial `gorm:"foreignKey:IndustrialID"`
}

func Migrate() {
	DATABASE.AutoMigrate(&model.User{})
	DATABASE.AutoMigrate(&model.Member{})
	DATABASE.AutoMigrate(&model.Companies{})
	DATABASE.AutoMigrate(&model.Industrial{})

	DATABASE.AutoMigrate(&MemberRelation{})
	DATABASE.AutoMigrate(&CompaniesRelation{})
}
