package database

import (
	"encoding/json"
	"first_go/src/lib"
	"first_go/src/model"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Seeders() {
	adminSeeder()
	mentorSeeder()
	subjectSeeder()
}

func adminSeeder() {
	email := "administrator@gmail.com"
	username := "administrator"
	password := "administrator"
	name := "Administrator"

	roles, _ := json.Marshal([]string{"admin"})
	provider, _ := json.Marshal([]string{"app"})
	hash, errHashing := bcrypt.GenerateFromPassword([]byte(password), 13)
	if errHashing != nil {
		panic("Failed To Hashing Seeder Addmin")
	}
	var vPassword string = string(hash)

	tx := DATABASE.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	user := model.User{
		Email:    email,
		Username: username,
		Password: &vPassword,
		Roles:    roles,
		Provider: provider,
	}

	admin := model.Admin{
		Name: name,
	}

	userAdmin := model.UserAdmin{
		User:  user,
		Admin: admin,
	}

	if err := tx.Create(&userAdmin).Error; err != nil {
		tx.Rollback()
		panic("Failed To Create Seeder Addmin")
	}

	tx.Commit()
	fmt.Println("Succes Create Admin Seeder")
}

var mentors = []struct {
	Email    string
	Username string
	Password string
	Name     string
}{
	{
		Email:    "bagus.yanuar613@gmail.com",
		Username: "bagus.yanuar613",
		Password: "bagus123",
		Name:     "Bagus Yanuar Arpri Dinata",
	},
	{
		Email:    "yulia.prastika@gmail.com",
		Username: "yulia.prastika",
		Password: "yulia123",
		Name:     "Yulia Prastika",
	},
	{
		Email:    "pradana.mahendra123@gmail.com",
		Username: "pradana.mahendra123",
		Password: "pradana123",
		Name:     "Pradana Mahendra",
	},
	{
		Email:    "taufiq.muhajir123@gmail.com",
		Username: "taufiq.muhajir123",
		Password: "taufiq123",
		Name:     "M Taufiq Muhajir",
	},
}

func mentorSeeder() {
	roles, _ := json.Marshal([]string{"mentor"})
	provider, _ := json.Marshal([]string{"app"})
	tx := DATABASE.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	for _, v := range mentors {

		hash, errHashing := bcrypt.GenerateFromPassword([]byte(v.Password), 13)
		if errHashing != nil {
			tx.Rollback()
			panic("Failed To Hashin Password")
		}
		var vPassword string = string(hash)
		user := model.User{
			Email:    v.Email,
			Username: v.Username,
			Password: &vPassword,
			Roles:    roles,
			Provider: provider,
		}

		mentor := model.Mentor{
			Name: v.Name,
		}
		userMentor := model.UserMentor{
			User:   user,
			Mentor: mentor,
		}

		if err := tx.Create(&userMentor).Error; err != nil {
			tx.Rollback()
			panic("Failed To Create Seeders")
		}
	}
	tx.Commit()
}

var subjects = []struct {
	Name string
}{
	{
		Name: "Matematika",
	},
	{
		Name: "Bahasa Indonesia",
	},
	{
		Name: "Bahasa Inggris",
	},
	{
		Name: "IPA",
	},
	{
		Name: "IPS",
	},
	{
		Name: "Komputer",
	},
}

func subjectSeeder() {
	tx := DATABASE.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	for _, v := range subjects {
		subject := model.Subject{
			Name: v.Name,
			Slug: lib.MakeSlug(v.Name),
		}

		if err := tx.Create(&subject).Error; err != nil {
			tx.Rollback()
			panic("Failed To Create Seeders")
		}
	}
	tx.Commit()
}
