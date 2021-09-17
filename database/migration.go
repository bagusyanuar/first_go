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

type JobsRelation struct {
	model.Jobs
	Companies  model.Companies  `gorm:"foreignKey:CompaniesID"`
	Profession model.Profession `gorm:"foreignKey:ProfessionID"`
}

func Migrate() {
	DATABASE.AutoMigrate(&model.User{})
	DATABASE.AutoMigrate(&model.Member{})
	DATABASE.AutoMigrate(&model.Companies{})
	DATABASE.AutoMigrate(&model.Industrial{})
	DATABASE.AutoMigrate(&model.Cities{})
	DATABASE.AutoMigrate(&model.Jobs{})
	DATABASE.Exec("ALTER TABLE `jobs` CHANGE `id` `id` CHAR(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL;")
	DATABASE.AutoMigrate(&model.Profession{})

	DATABASE.AutoMigrate(&MemberRelation{})
	DATABASE.AutoMigrate(&CompaniesRelation{})
	DATABASE.AutoMigrate(&JobsRelation{})
}
