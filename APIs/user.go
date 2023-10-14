package apis

import (
	"encoding/json"
	"log"
	"net/http"
	"rto/database"
	"rto/logic"
	structs "rto/struct"
	"time"

	"github.com/google/uuid"
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

func CreateRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Create request API called")
	now := time.Now()
	var test string
	test = "Report submitted successfully"
	var ByRTO, Social bool
	var requestData structs.Report

	requestData.ReportId = uuid.New().String()
	requestData.RTOApproved = r.FormValue("ByRto")
	requestData.Comments = r.FormValue("Comment")
	requestData.ImageUrls = r.FormValue("ImageURL")
	requestData.Locations = r.FormValue("Locations")
	requestData.Offense = r.FormValue("Offense")
	requestData.ReportedBy = r.FormValue("ReportedBy")
	requestData.Social = r.FormValue("Social")
	requestData.VehicleNumber = r.FormValue("VehicleNumber")

	if requestData.RTOApproved == "True" {
		ByRTO = true
	}

	if requestData.Social == "True" {
		Social = true
	}

	requestData.TotalFine = 100

	reportId, err := database.SubmitReport(requestData, ByRTO, Social)
	if err != nil {
		test = "Failed to create request please try again later !"
	}

	test += " : " + reportId
	log.Println("Create request API successfully completed in " + time.Since(now).String())
	json.NewEncoder(w).Encode(test)
}
