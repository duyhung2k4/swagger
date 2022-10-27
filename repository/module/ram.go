package module

import (
	connect "app/config"
	"app/repository/model"
)

func GetRam() []model.Ram {

	db := connect.Connect()

	var rams []model.Ram

	db.Find(&rams)

	return rams
}

func AddRam(ram model.Ram) {

	db := connect.Connect()

	db.Create(&ram)
}

func UpdateRam(ram model.Ram) {

	db := connect.Connect()

	db.Table("rams").Where("id = ?", ram.Id).Updates(&ram)
}

func DeleteRam(ram model.Ram) {

	db := connect.Connect()

	db.Delete(&ram)
}
