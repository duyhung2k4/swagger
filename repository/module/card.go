package module

import (
	connect "app/config"
	"app/repository/model"
)

func AddCard(data model.Card) error {

	db := connect.Connect()

	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func GetCard() []model.Card {

	db := connect.Connect()

	var cards []model.Card

	db.Find(&cards)

	return cards

}

func UpdateCard(data model.Card) {

	db := connect.Connect()

	db.Table("cards").Where("id = ?", data.Id).Updates(&data)

}

func DeleteCard(card model.Card) {

	db := connect.Connect()

	db.Delete(&card)
}
