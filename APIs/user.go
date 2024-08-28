package apis

import (
	"encoding/json"
	"log"
	"net/http"
	"rto/database"
	"rto/logic"
	"strings"
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
	var response structs.ReportResopnse
	test = "Report submitted successfully"
	var ByRTO, Social bool
	var requestData structs.Report
	go database.Logs("offense form value", r.FormValue("Offense"))
	requestData.ReportId = uuid.New().String()
	requestData.RTOApproved = r.FormValue("ByRto")
	requestData.Comments = r.FormValue("Comment")
	requestData.ImageUrls = r.FormValue("ImageURL")
	requestData.Locations = r.FormValue("Locations")
	requestData.Offense = r.FormValue("Offense")
	requestData.ReportedBy = r.FormValue("ReportedBy")
	requestData.Social = r.FormValue("Social")
	requestData.VehicleNumber = r.FormValue("VehicleNumber")
	go database.Logs("offense val", requestData.Offense)

	if requestData.RTOApproved == "True" {
		ByRTO = true
	}

	if requestData.Social == "True" {
		Social = true
	}

	offenceSlice := formatSlice(requestData.Offense)
	requestData.TotalFine = checkFineForOffence(offenceSlice)

	log.Println(requestData.TotalFine, "<- Fined for ")

	reportId, err := database.SubmitReport(requestData, ByRTO, Social)
	if err != nil {
		test = "Failed to create request please try again later !"
		go database.Logs("Debug", test)
	}
	response.RequestId = reportId
	response.TotalFine = requestData.TotalFine
	l1 := "Create request API successfully completed in " + time.Since(now).String()
	go database.Logs("Debug", l1)
	json.NewEncoder(w).Encode(response)
}

func checkFineForOffence(offence []string)int{
	fine := make(map[string]int)
	fine["Overspeeding"] = 5000
	fine["Drink and drive offence"] = 10000
	fine["Riding without wearing a helmet (rider/pillion rider)"] = 1000
	fine["Driving without wearing a seatbelt"] = 1000
	fine["Using a mobile phone while driving/riding"] = 5000
	fine["Overloading of two-wheeler (triple riding)"] = 1000
	fine["Offence related to air/noise pollution"] = 1000
	fine["Driving/riding without number plate"] = 500
	fine["Minor riding or driving a vehicle"] = 25000
	fine["Lane discipline offences"] = 500
	fine["No-parking offences"] = 500
	fine["One way offence"] = 500
	fine["Offence related to alteration of vehicle"] = 5000

	var	totalFine int
	for _, o := range offence {
		totalFine += fine[o]
	}

	return totalFine
}

func formatSlice(offence string)[]string{
	l1 := strings.Trim(offence, "[")
	l2 := strings.Trim(l1, "]")
	l3 := strings.Replace(l2, `"`, "", -1)
	k := strings.Split(l3, ",")

	return k
}

func SignUpUser(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	msg := "User Created successfully"
	var request structs.UserRequest

	request.UserName = r.FormValue("UserName")
	request.FirstName = r.FormValue("FirstName")
	request.LastName = r.FormValue("LastName")
	request.Email = r.FormValue("Email")
	request.Password = r.FormValue("Password")

	pass, err := logic.Encrypt(request.Password)
	if err != nil {
		log.Println(err.Error(), " <- Error")
	}

	err2 := database.Upsertuser(request, pass)
	if err2 != nil {
		log.Println(err.Error(), "<- Error")
	}

	json.NewEncoder(w).Encode(msg)
	log.Println("SignUp API called completed in : ", time.Since(now))
}
