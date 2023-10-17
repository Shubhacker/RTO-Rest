package apis

import (
	"encoding/json"
	"log"
	"net/http"
	"rto/database"
	"time"
)

func SocialFetch(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	data, _, err := database.FetchSocialFromDB("")
	if err != nil {
		log.Println(err.Error(), " <- Error")
	}

	json.NewEncoder(w).Encode(data)

	log.Println("Social API call : ", time.Since(now))
}
