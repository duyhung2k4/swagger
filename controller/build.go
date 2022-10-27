package controller

import (
	"app/repository/model"
	module "app/repository/module"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/go-chi/jwtauth/v5"
)

func GetBuild(w http.ResponseWriter, r *http.Request) {
	log.Println("X")
	_, claims, _ := jwtauth.FromContext(r.Context())

	fmt.Println("type: ", reflect.TypeOf(claims["id"]))

	var id int = int(claims["id"].(float64))

	result, _ := json.Marshal(module.GetBuild(id))

	w.Write(result)

}

func AddBuild(w http.ResponseWriter, r *http.Request) {

	var pc model.Build

	json.NewDecoder(r.Body).Decode(&pc)

	_, claims, _ := jwtauth.FromContext(r.Context())

	user_id := int(claims["id"].(float64))

	pc.UserId = user_id

	fmt.Println("Pc: ", pc)

	err := module.AddBuild(pc)

	if err != nil {

		result, _ := json.Marshal(map[string]interface{}{
			"warn": "error",
		})

		w.Write(result)

	} else {
		result, _ := json.Marshal(map[string]interface{}{
			"text": "Done",
		})

		w.Write(result)

	}

}

func DeleteBuild(w http.ResponseWriter, r *http.Request) {

	_, claims, _ := jwtauth.FromContext(r.Context())

	var pc model.Build

	json.NewDecoder(r.Body).Decode(&pc)

	user_id := int(claims["id"].(float64))

	module.DeleteBuild(user_id, pc.Id)

	result, _ := json.Marshal(map[string]interface{}{
		"text": "Done",
	})

	w.Write(result)

}
