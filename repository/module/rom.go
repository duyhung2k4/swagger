package module

import (
	connect "app/config"
	"app/repository/model"
)

func GetRom() []model.Rom {

	db := connect.Connect()

	var roms []model.Rom

	db.Find(&roms)

	return roms

}

func AddRom(rom model.Rom) {

	db := connect.Connect()

	db.Create(&rom)
}

func UpdateRom(rom model.Rom) {

	db := connect.Connect()

	db.Table("roms").Where("id = ?", rom.Id).Updates(&rom)
}

func DeleteRom(rom model.Rom) {

	db := connect.Connect()

	db.Delete(&rom)
}
