package module

import (
	connect "app/config"
	"app/repository/model"
	"fmt"
)

func GetBuild(id int) []model.Build {
	db := connect.Connect()
	var builds []model.Build

	fmt.Println("User_id: ", id)

	db.Where("user_id = ?", id).Model(&model.Build{}).
		Preload("User").
		Preload("Card").
		Preload("Main").
		Preload("Cpu").
		Preload("Ram").
		Preload("Rom").
		Preload("Psup").
		Preload("HeatDissipation").Find(&builds)

	return builds
}

func AddBuild(pc model.Build) error {

	db := connect.Connect()

	if err := db.Create(&pc).Error; err != nil {
		return err
	}

	return nil
}

func DeleteBuild(user_id int, id int) {

	db := connect.Connect()

	db.Where("user_id = ? AND id = ?", user_id, id).Delete(&model.Build{})

}
