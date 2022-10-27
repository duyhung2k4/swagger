package controller

import (
	"app/repository/model"
	"app/repository/module"
	"encoding/json"
	"net/http"
)

func GetPSup(w http.ResponseWriter, r *http.Request) {

	result, _ := json.Marshal(module.GetPsup())

	w.Write(result)
}

func AddPsup(w http.ResponseWriter, r *http.Request) {

	var psup model.Psup

	json.NewDecoder(r.Body).Decode(&psup)

	module.AddPsup(psup)

	result, _ := json.Marshal(map[string]interface{}{
		"text": "Done",
	})

	w.Write(result)
}

func UpdatePsup(w http.ResponseWriter, r *http.Request) {

	var psup model.Psup

	json.NewDecoder(r.Body).Decode(&psup)

	module.UpdatePsup(psup)

	result, _ := json.Marshal(map[string]interface{}{
		"text": "Done",
	})

	w.Write(result)
}

func DeletePsup(w http.ResponseWriter, r *http.Request) {

	var psup model.Psup

	json.NewDecoder(r.Body).Decode(&psup)

	module.DeletePsup(psup)

	result, _ := json.Marshal(map[string]interface{}{
		"text": "Delete Psup",
	})

	w.Write(result)
}
