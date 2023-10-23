package structs

type AddComment struct {
	RequestId string
	Comment   string
	CommentBy string
}

type CommentData struct {
	ReportId       string
	Comment        string
	CommentBy      string
	ReportLikes    int
	ReportDisLikes int
	CreatedAt      string
}
