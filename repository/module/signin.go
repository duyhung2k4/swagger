package module

import (
	connect "app/config"
	"app/repository/model"
	"fmt"

	"gorm.io/gorm"
)

func CheckAccoutSignIn(name string) model.Accout {

	db := connect.Connect()

	var accout []model.Accout

	db.Where("name = ?", name).Find(&accout)

	if len(accout) == 0 {

		user := model.Accout{Id: 0, Name: "", Pass: "", Email: ""}

		return user
	}

	return accout[0]
}

func CreateAccout(user model.Accout) *gorm.DB {

	db := connect.Connect()

	fmt.Println("accout: ", user)

	result := db.Create(&user)

	return result
}
