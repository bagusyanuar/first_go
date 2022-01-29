package database

import "first_go/src/model"

type MemberRelation struct {
	model.Member
	User model.User `gorm:"foreignKey:UserID"`
}

type MentorRelation struct {
	model.Mentor
	User model.User `gorm:"foreignKey:UserID"`
}

type AdminRelation struct {
	model.Admin
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
	DATABASE.AutoMigrate(&model.Admin{})
	DATABASE.AutoMigrate(&model.Companies{})
	DATABASE.AutoMigrate(&model.Industrial{})
	DATABASE.AutoMigrate(&model.Cities{})
	DATABASE.AutoMigrate(&model.Jobs{})
	DATABASE.Exec("ALTER TABLE `jobs` CHANGE `id` `id` CHAR(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL;")
	DATABASE.AutoMigrate(&model.Profession{})
	DATABASE.AutoMigrate(&model.Subject{})
	DATABASE.Exec("ALTER TABLE `subjects` CHANGE `icon` `icon` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL")
	DATABASE.Exec("ALTER TABLE `subjects` CHANGE `id` `id` INT(11) NOT NULL AUTO_INCREMENT;")

	DATABASE.AutoMigrate(&MemberRelation{})
	DATABASE.AutoMigrate(&MentorRelation{})
	DATABASE.AutoMigrate(&AdminRelation{})
	DATABASE.AutoMigrate(&CompaniesRelation{})
	DATABASE.AutoMigrate(&JobsRelation{})
}
