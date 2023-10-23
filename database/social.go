package database

import (
	"log"
	structs "rto/struct"
)

func FetchSocialFromDB(requestBy string) (structs.SocialResponse, []structs.SocialData, []string, error) {
	var response []structs.SocialData
	var responseD structs.SocialResponse

	var requestIdSlice []string
	query := `select report_id ,image_urls ,"location" ,offense ,is_submitted_by_rto ,total_fine ,rto_approved from report.public_report order by RANDOM() limit 10;`

	rows, err := DB.Query(query)
	if err != nil {
		log.Println("Issue (FetchSocialFromDB) : ", err.Error())
		return responseD, response, requestIdSlice, err
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
	responseD.Social = response
	return responseD, response, requestIdSlice, nil
}

func AddComment(data structs.AddComment) error {
	var inputArgs []interface{}
	var reportId string
	query := `insert into report.report_comment (report_id, "comment", comment_by) values ($1, $2, $3) returning report_id`
	inputArgs = append(inputArgs, data.RequestId)
	inputArgs = append(inputArgs, data.Comment)
	inputArgs = append(inputArgs, data.CommentBy)

	err2 := DB.QueryRow(query, inputArgs...).Scan(&reportId)
	if err2 != nil {
		return err2
	}

	return nil
}

func FetchComments(requestId string) ([]structs.CommentData, error) {
	var commentData []structs.CommentData

	query := `select report_id, "comment", comment_by, report_likes, report_dislike, created_at from report.report_comment where report_id = $1`

	rows, err := DB.Query(query, requestId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comment structs.CommentData
		rows.Scan(
			&comment.ReportId,
			&comment.Comment,
			&comment.CommentBy,
			&comment.ReportLikes,
			&comment.ReportDisLikes,
			&comment.CreatedAt,
		)
	}

	return commentData, nil
}
