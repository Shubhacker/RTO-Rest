package database

import (
	"log"
	structs "rto/struct"
)

func FetchSocialFromDB(requestBy string) ([]structs.SocialData, []string, error) {
	var response []structs.SocialData
	var requestIdSlice []string
	query := `select report_id ,image_urls ,"location" ,offense ,is_submitted_by_rto ,total_fine ,rto_approved from report.public_report order by RANDOM() limit 10;`

	rows, err := DB.Query(query)
	if err != nil {
		log.Println("Issue (FetchSocialFromDB) : ", err.Error())
		return response, requestIdSlice, err
	}
	for rows.Next() {
		var DBdata structs.SocialData
		rows.Scan(
			&DBdata.RequestId,
			&DBdata.ImageUrls,
			&DBdata.Location,
			&DBdata.Offence,
			&DBdata.SubmittedByRTO,
			&DBdata.TotalFine,
			&DBdata.RTOApproved,
		)
		response = append(response, DBdata)
		requestIdSlice = append(requestIdSlice, DBdata.RequestId)
	}

	return response, requestIdSlice, nil
}
