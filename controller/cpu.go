package controller

import (
	"app/repository/model"
	"app/repository/module"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetCpu(w http.ResponseWriter, r *http.Request) {

	result, _ := json.Marshal(module.GetCpu())

	w.Write(result)

}

func AddCpu(w http.ResponseWriter, r *http.Request) {

	var cpu model.Cpu

	json.NewDecoder(r.Body).Decode(&cpu)

	module.AddCpu(cpu)

	result, _ := json.Marshal(map[string]interface{}{
		"text": "Done",
	})

	w.Write(result)
}

func UpdateCpu(w http.ResponseWriter, r *http.Request) {

	var cpu model.Cpu

	json.NewDecoder(r.Body).Decode(&cpu)

	module.UpdateCpu(cpu)

	result, _ := json.Marshal(map[string]interface{}{
		"text": "Done",
	})

	w.Write(result)
}

func DeleteCpu(w http.ResponseWriter, r *http.Request) {

	var cpu model.Cpu

	json.NewDecoder(r.Body).Decode(&cpu)

	fmt.Println(cpu)

	module.DeleteCpu(cpu)

	result, _ := json.Marshal(map[string]interface{}{
		"text": "Done",
	})

	w.Write(result)
}
