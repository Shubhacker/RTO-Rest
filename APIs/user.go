package apis

import (
	"encoding/json"
	"log"
	"net/http"
	"rto/database"
	"rto/logic"
	structs "rto/struct"
)

func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside Login API")
	var responseData structs.LoginResponse
	var userData structs.User
	var authenticated bool
	var msg string

	username := r.FormValue("username")
	password := r.FormValue("password")
	log.Println(username)
	log.Println(password)

	encryptPass, _ := logic.Encrypt(password)

	user, err := database.FetchUserInfo(username)
	if err != nil {
		log.Println(err.Error())
	}

	if user.Password == encryptPass {
		authenticated = true
	}
	responseData.Authenticated = authenticated

	if !authenticated {
		msg = "Wrong Password"
		responseData.Message = msg
		// return &responseData, nil
	}
	msg = "Logic Success"

	userData.FirstName = user.FirstName
	userData.LastName = user.LastName
	userData.UserId = user.UserId
	userData.Username = user.Username
	JWT, err := logic.GenerateToken(&user)
	if err != nil {
		// return &response, err
	}
	log.Println(*JWT)
	userData.JWT = *JWT

	responseData.Message = msg
	responseData.UserInfo = userData

	json.NewEncoder(w).Encode(responseData)

}
