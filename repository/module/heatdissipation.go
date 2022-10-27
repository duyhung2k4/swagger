package module

import (
	connect "app/config"
	"app/repository/model"
)

func GetHeatDissipation() []model.HeatDissipation {

	db := connect.Connect()

	var heat_dissipations []model.HeatDissipation

	db.Find(&heat_dissipations)

	return heat_dissipations

}

func AddHeatDissipation(heat_dissipation model.HeatDissipation) {

	db := connect.Connect()

	db.Create(&heat_dissipation)
}

func UpdateHeatDissipation(heat_dissipation model.HeatDissipation) {

	db := connect.Connect()

	db.Table("heat_dissipations").Where("id = ?", heat_dissipation.Id).Updates(&heat_dissipation)
}

func DeleteHeatDissipation(heat_dissipation model.HeatDissipation) {

	db := connect.Connect()

	db.Delete(&heat_dissipation)
}
