package controller

import (
	"app/repository/model"
	"app/repository/module"
	"encoding/json"
	"net/http"
)

func GetMain(w http.ResponseWriter, r *http.Request) {

	result, _ := json.Marshal(module.GetMain())

	w.Write(result)

}

func AddMain(w http.ResponseWriter, r *http.Request) {

	var main model.Main

	json.NewDecoder(r.Body).Decode(&main)

	module.AddMain(main)

	result, _ := json.Marshal(map[string]interface{}{
		"text": "Done",
	})

	w.Write(result)
}

func UpdateMain(w http.ResponseWriter, r *http.Request) {

	var main model.Main

	json.NewDecoder(r.Body).Decode(&main)

	module.UpdateMain(main)

	result, _ := json.Marshal(map[string]interface{}{
		"text": "Done",
	})

	w.Write(result)
}

func DeleteMain(w http.ResponseWriter, r *http.Request) {

	var main model.Main

	json.NewDecoder(r.Body).Decode(&main)

	module.DeleteMain(main)

	result, _ := json.Marshal(map[string]interface{}{
		"text": "Done",
	})

	w.Write(result)
}
