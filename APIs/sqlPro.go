package apis

import (
	"encoding/json"
	"log"
	"net/http"
	"rto/logic"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

// Not belong to these project
func Mergequery(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	sql := r.FormValue("query")
	inputArgs := r.FormValue("input")

	//converting inputargs from string to slice
	slice := strings.Fields(inputArgs)
	log.Println(inputArgs, " <= InputArgs converted")

	// converting ? to $ for checking number of input parameter required
	sqlval := sqlx.Rebind(sqlx.DOLLAR, sql)
	log.Println(sqlval, "Converted to $")
	que := logic.ReplaceDollarWithData(sqlval, slice)
	log.Println(que, "<- query replace $")
	log.Println(time.Since(now), " <- Time took to provide  responce")
	json.NewEncoder(w).Encode(que)
}
