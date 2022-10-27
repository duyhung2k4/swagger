package controller

import (
	"app/repository/model"
	"app/repository/module"
	"encoding/json"
	"fmt"
	"net/http"
)

func AddCard(w http.ResponseWriter, r *http.Request) {
	var card model.Card

	json.NewDecoder(r.Body).Decode(&card)

	fmt.Println(card)

	module.AddCard(card)

	result, _ := json.Marshal(map[string]interface{}{
		"status": "Done",
	})

	w.Write(result)
}

func GetCard(w http.ResponseWriter, r *http.Request) {

	//_, claims, _ := jwtauth.FromContext(r.Context())

	result, _ := json.Marshal(module.GetCard())

	w.Write(result)

}

func UpdateCard(w http.ResponseWriter, r *http.Request) {

	var card model.Card

	json.NewDecoder(r.Body).Decode(&card)

	module.UpdateCard(card)

	result, _ := json.Marshal(map[string]interface{}{
		"text": "Done",
	})

	w.Write(result)

}

func DeleteCard(w http.ResponseWriter, r *http.Request) {

	var card model.Card

	json.NewDecoder(r.Body).Decode(&card)

	module.DeleteCard(card)

	result, _ := json.Marshal(map[string]interface{}{
		"text": "Done",
	})

	w.Write(result)
}
