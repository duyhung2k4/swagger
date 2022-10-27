package controller

import (
	"app/repository/model"
	"app/repository/module"
	"encoding/json"
	"net/http"
)

func GetRom(w http.ResponseWriter, r *http.Request) {

	result, _ := json.Marshal(module.GetRom())

	w.Write(result)
}

func AddRom(w http.ResponseWriter, r *http.Request) {

	var rom model.Rom

	json.NewDecoder(r.Body).Decode(&rom)

	module.AddRom(rom)

	result, _ := json.Marshal(map[string]interface{}{
		"text": "Done",
	})

	w.Write(result)
}

func UpdateRom(w http.ResponseWriter, r *http.Request) {

	var rom model.Rom

	json.NewDecoder(r.Body).Decode(&rom)

	module.UpdateRom(rom)

	result, _ := json.Marshal(map[string]interface{}{
		"text": "Done",
	})

	w.Write(result)
}

func DeleteRom(w http.ResponseWriter, r *http.Request) {

	var rom model.Rom

	json.NewDecoder(r.Body).Decode(&rom)

	module.DeleteRom(rom)

	result, _ := json.Marshal(map[string]interface{}{
		"text": "Done",
	})

	w.Write(result)
}
