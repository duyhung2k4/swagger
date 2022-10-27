package module

import (
	connect "app/config"
	"app/repository/model"
)

func GetCpu() []model.Cpu {

	db := connect.Connect()

	var cpus []model.Cpu

	db.Find(&cpus)

	return cpus
}

func AddCpu(cpu model.Cpu) {

	db := connect.Connect()

	db.Create(&cpu)
}

func UpdateCpu(cpu model.Cpu) {

	db := connect.Connect()

	db.Table("cpus").Where("id = ?", cpu.Id).Updates(&cpu)
}

func DeleteCpu(cpu model.Cpu) {

	db := connect.Connect()
	db.Delete(&cpu)
}
