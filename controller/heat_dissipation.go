package controller

import (
	"app/repository/model"
	"app/repository/module"
	"encoding/json"
	"net/http"
)

func GetHeatDissipation(w http.ResponseWriter, r *http.Request) {

	result, _ := json.Marshal(module.GetHeatDissipation())

	w.Write(result)
}

func AddHeatDissipation(w http.ResponseWriter, r *http.Request) {

	var heat_dissipation model.HeatDissipation

	json.NewDecoder(r.Body).Decode(&heat_dissipation)

	module.AddHeatDissipation(heat_dissipation)

	result, _ := json.Marshal(map[string]interface{}{
		"text": "Done",
	})

	w.Write(result)
}

func UpdateHeatDissipation(w http.ResponseWriter, r *http.Request) {

	var heat_dissipation model.HeatDissipation

	json.NewDecoder(r.Body).Decode(&heat_dissipation)

	module.UpdateHeatDissipation(heat_dissipation)

	result, _ := json.Marshal(map[string]interface{}{
		"text": "Done",
	})

	w.Write(result)
}

func DeleteHeatDissipation(w http.ResponseWriter, r *http.Request) {

	var heat_dissipation model.HeatDissipation

	json.NewDecoder(r.Body).Decode(&heat_dissipation)

	module.DeleteHeatDissipation(heat_dissipation)

	result, _ := json.Marshal(map[string]interface{}{
		"text": "Delete",
	})

	w.Write(result)
}
