package apis

import "github.com/gorilla/mux"

func Apis() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/login/", Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/createrequest/", CreateRequest).Methods("POST", "OPTIONS")
	r.HandleFunc("/signUp/", SignUpUser).Methods("POST", "OPTIONS")
	return r
}
