package module

import (
	connect "app/config"
	"app/repository/model"
)

func CheckAccout(name string, pass string) model.Accout {

	db := connect.Connect()

	var accout []model.Accout

	db.Where("name = ? AND pass = ?", name, pass).Find(&accout)

	if len(accout) == 0 {

		user := model.Accout{Id: 0, Name: "", Pass: "", Email: ""}

		return user
	}

	return accout[0]
}
