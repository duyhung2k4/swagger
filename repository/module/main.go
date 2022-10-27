package module

import (
	connect "app/config"
	"app/repository/model"
)

func GetMain() []model.Main {

	db := connect.Connect()

	var mains []model.Main

	db.Find(&mains)

	return mains

}

func AddMain(main model.Main) {

	db := connect.Connect()

	db.Create(&main)
}

func UpdateMain(main model.Main) {

	db := connect.Connect()

	db.Table("mains").Where("id = ? ", main.Id).Updates(&main)
}

func DeleteMain(main model.Main) {

	db := connect.Connect()

	db.Delete(&main)
}
