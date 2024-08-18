package apis

import "github.com/gorilla/mux"

func Apis() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/login/", Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/createrequest/", CreateRequest).Methods("POST", "OPTIONS")
	r.HandleFunc("/signUp/", SignUpUser).Methods("POST", "OPTIONS")
	r.HandleFunc("/social/", SocialFetch).Methods("GET", "OPTIONS")
	r.HandleFunc("/getMyRequest/", FetchMyRequest).Methods("GET", "OPTIONS")
	// Implementing
	r.HandleFunc("/", FetchComments).Methods("GET", "OPTIONS")
	// To implement
	r.HandleFunc("/FetchComments/", FetchComments).Methods("GET", "OPTIONS")
	r.HandleFunc("/AddComment/", AddComment).Methods("GET", "OPTIONS")
	r.HandleFunc("/LikeDislikeComment/", LikeDislikeComment).Methods("GET", "OPTIONS")
	r.HandleFunc("/MergeQuery/", Mergequery).Methods("GET", "OPTIONS")

	return r
}
