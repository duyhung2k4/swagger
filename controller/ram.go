package controller

import (
	"app/repository/model"
	"app/repository/module"
	"encoding/json"
	"net/http"
)

func GetRam(w http.ResponseWriter, R *http.Request) {

	result, _ := json.Marshal(module.GetRam())

	w.Write(result)

}

func AddRam(w http.ResponseWriter, r *http.Request) {

	var ram model.Ram

	json.NewDecoder(r.Body).Decode(&ram)

	module.AddRam(ram)

	result, _ := json.Marshal(map[string]interface{}{
		"text": "Done",
	})

	w.Write(result)
}

func UpdateRam(w http.ResponseWriter, r *http.Request) {

	var ram model.Ram

	json.NewDecoder(r.Body).Decode(&ram)

	module.UpdateRam(ram)

	result, _ := json.Marshal(map[string]interface{}{
		"text": "Done",
	})

	w.Write(result)
}

func DeleteRam(w http.ResponseWriter, r *http.Request) {

	var ram model.Ram

	json.NewDecoder(r.Body).Decode(&ram)

	module.DeleteRam(ram)

	result, _ := json.Marshal(map[string]interface{}{
		"text": "Done",
	})

	w.Write(result)
}
