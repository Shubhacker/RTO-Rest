package apis

import (
	"encoding/json"
	"net/http"
	"rto/database"
	structs "rto/struct"
	"time"
)

func CreateUserPaymentInfo(w http.ResponseWriter, r *http.Request){
	now := time.Now()
	var upf structs.UserPaymentsInfo
	var userPaymentId string
	var err error

	upf.UserName = r.FormValue("username")
	upf.UserId = r.FormValue("userId")
	upf.UPIID = r.FormValue("upiId")

	// Check if UPI ID/username/userID already present in DB
	isPresent := database.IsUserPaymentInfoPresent(upf.UserId)

	if isPresent{
		// Update value Pending
	} else{
		// Insert value
		userPaymentId, err = database.AddUserPaymentInformation(upf)
		if err != nil {
			go database.Logs("Error CreateUserPaymentInfo: ", err.Error())
		}
	}
	var res structs.UserPaymentInfoResponse
	res.ID = userPaymentId

	go database.Logs("Time took for CreateUserPaymentInfo: ", time.Since(now).String())
	json.NewEncoder(w).Encode(res)
}