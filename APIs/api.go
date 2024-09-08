package apis

import "github.com/gorilla/mux"

func Apis() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/login/", Login).Methods("POST", "OPTIONS")
	r.HandleFunc("/createrequest/", CreateRequest).Methods("POST", "OPTIONS")
	r.HandleFunc("/signUp/", SignUpUser).Methods("POST", "OPTIONS")
	r.HandleFunc("/social/", SocialFetch).Methods("GET", "OPTIONS")
	r.HandleFunc("/getMyRequest/", FetchMyRequest).Methods("GET", "OPTIONS")
	r.HandleFunc("/FetchComments/", FetchComments).Methods("GET", "OPTIONS")
	r.HandleFunc("/AddComment/", AddComment).Methods("GET", "OPTIONS")
	r.HandleFunc("/LikeDislikeComment/", AddLikesDisLikes).Methods("POST", "OPTIONS")
	// Implementing

	// To implement

	// Not belong to these project
	r.HandleFunc("/MergeQuery/", Mergequery).Methods("GET", "OPTIONS")

	return r
}
