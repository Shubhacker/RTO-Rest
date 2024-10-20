package apis

import (
	"github.com/gorilla/mux"

)

func Apis() *mux.Router {
	r := mux.NewRouter()
	// User APIs
	r.HandleFunc("/login/", Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/signUp/", SignUpUser).Methods("POST", "OPTIONS")

	// Report APIs
	r.HandleFunc("/createrequest/", CreateRequest).Methods("POST", "OPTIONS")
	r.HandleFunc("/social/", SocialFetch).Methods("GET", "OPTIONS")
	r.HandleFunc("/getMyRequest/", FetchMyRequest).Methods("GET", "OPTIONS")

	// Comment APIs
	r.HandleFunc("/FetchComments/", FetchComments).Methods("GET", "OPTIONS")
	r.HandleFunc("/AddComment/", AddComment).Methods("GET", "OPTIONS")
	r.HandleFunc("/LikeDislikeComment/", AddLikesDisLikes).Methods("POST", "OPTIONS")
	
	// Payment APIs
	r.HandleFunc("/userPaymentInfo/", CreateUserPaymentInfo).Methods("POST", "OPTIONS")
	// r.HandleFunc("/paymentInfo/", CreateUserPaymentInfo).Methods("POST", "OPTIONS")

	r.HandleFunc("/checkChallan/", CheckChallan).Methods("POST", "OPTIONS")
	r.HandleFunc("/vehicleInfo/", VehicleInformation).Methods("POST", "OPTIONS")


	// Not belong to these project
	r.HandleFunc("/MergeQuery/", Mergequery).Methods("GET", "OPTIONS")

	return r
}
