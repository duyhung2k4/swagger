package controller

import (
	"app/repository/model"
	"app/repository/module"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var token *jwtauth.JWTAuth

	var user model.UserLogin

	json.NewDecoder(r.Body).Decode(&user)

	fmt.Println("User: ", user)

	accout := module.CheckAccout(user.Name, user.Pass)

	fmt.Println("Acc: ", accout)

	if accout.Id == 0 {
		result, _ := json.Marshal(map[string]interface{}{
			"text": "None exist",
		})

		fmt.Println(token)

		w.Write(result)
	} else {

		token := jwtauth.New("HS256", []byte("token"), nil)

		_, tokenString, _ := token.Encode(map[string]interface{}{
			"name": accout.Name,
			"id":   accout.Id,
		})

		fmt.Println("Token: ", tokenString)

		cookie := http.Cookie{
			Name:     "jwt",
			Value:    tokenString,
			HttpOnly: true,
			MaxAge:   60 * 60,
		}

		http.SetCookie(w, &cookie)

		result, _ := json.Marshal(map[string]interface{}{
			"token": tokenString,
		})

		w.Write(result)

	}

}
