package module

import (
	connect "app/config"
	"app/repository/model"
)

func GetPsup() []model.Psup {

	db := connect.Connect()

	var psups []model.Psup

	db.Find(&psups)

	return psups
}

func AddPsup(psup model.Psup) {

	db := connect.Connect()

	db.Create(&psup)
}

func UpdatePsup(psup model.Psup) {

	db := connect.Connect()

	db.Table("psups").Where("id = ?", psup.Id).Updates(&psup)
}

func DeletePsup(psup model.Psup) {

	db := connect.Connect()

	db.Delete(&psup)
}
