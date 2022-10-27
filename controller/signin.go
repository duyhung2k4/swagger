package controller

import (
	"app/repository/model"
	"app/repository/module"
	"encoding/json"
	"fmt"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {

	var newUser model.Accout

	json.NewDecoder(r.Body).Decode(&newUser)

	accout := module.CheckAccoutSignIn(newUser.Name)

	if accout.Id == 0 {
		user := model.Accout{Name: newUser.Name, Pass: newUser.Pass, Email: newUser.Email}

		fmt.Println("User: ", user)

		result := module.CreateAccout(user)

		if result.Error != nil {

			request, _ := json.Marshal(map[string]interface{}{
				"text": "Error",
			})

			w.Write(request)
		} else {
			request, _ := json.Marshal(map[string]interface{}{
				"text": "Done",
			})

			w.Write(request)
		}

	} else {

		result, _ := json.Marshal(map[string]interface{}{
			"Text": "Warning Exist",
		})

		w.Write(result)
	}

}
