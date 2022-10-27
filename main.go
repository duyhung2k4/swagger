package main

import (
	connect "app/config"
	"app/repository/model"
	router "app/routers"
)

func main() {

	db := connect.Connect()

	db.AutoMigrate(
		&model.Card{},
		&model.Main{},
		&model.Ram{},
		&model.Rom{},
		&model.Cpu{},
		&model.Accout{},
		&model.Psup{},
		&model.HeatDissipation{},
		&model.Build{},
	)

	router.Router()
}
