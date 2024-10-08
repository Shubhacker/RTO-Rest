package apis

import (
	"encoding/json"
	"log"
	"net/http"
	"rto/database"
	structs "rto/struct"
	"time"
)

func SocialFetch(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	data2, _, _, err := database.FetchSocialFromDB("")
	if err != nil {
		log.Println(err.Error(), " <- Error")
	}

	json.NewEncoder(w).Encode(data2)

	log.Println("Social API call : ", time.Since(now))
}

func FetchMyRequest(w http.ResponseWriter, r *http.Request){
	log.Println("FetchMyRequest Called")
	now := time.Now()
	submittedBy := r.FormValue("submittedBy")
	data2, _, _, err := database.FetchSocialFromDB(submittedBy)
	if err != nil {
		log.Println(err.Error(), " <- Error")
	}

	json.NewEncoder(w).Encode(data2)

	log.Println("FetchMyRequest API call : ", time.Since(now))
}

func FetchComments(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	RequestId := r.FormValue("RequestId")
	var commentRes structs.CommentResponse

	commentData, err := database.FetchComments(RequestId)
	if err != nil {
		log.Println(err.Error(), " <- Error returned from DB")
	}
	commentRes.Comments = commentData

	json.NewEncoder(w).Encode(commentRes)

	log.Println("FetchComments API call : ", time.Since(now))
}

func AddComment(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	var commentData structs.AddComment
	commentData.RequestId = r.FormValue("RequestId")
	commentData.CommentBy = r.FormValue("CommentBy")
	commentData.Comment = r.FormValue("Comment")

	err := database.AddComment(commentData)
	if err != nil {
		log.Println(err.Error(), "<- Error returned from DB")
	}

	log.Println("AddComment API call : ", time.Since(now))
}

func LikeDislikeComment(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	log.Println("LikeDislikeComment API call : ", time.Since(now))
}
