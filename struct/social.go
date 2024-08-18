package structs

type SocialResponse struct {
	Social []SocialData
}

type SocialData struct {
	RequestId      string
	ImageUrls      string
	Location       string
	Offence        string
	SubmittedByRTO bool
	TotalFine      int
	RTOApproved    bool
	Submitted_at string
}
