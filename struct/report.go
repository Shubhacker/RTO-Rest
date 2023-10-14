package structs

type ReportCreate struct {
	ImageURL      []string `json:"imageUrl,omitempty"`
	Locations     string   `json:"locations,omitempty"`
	Offense       []string `json:"offense,omitempty"`
	ByRto         string   `json:"ByRTO,omitempty"`
	Social        string   `json:"Social,omitempty"`
	Comment       string   `json:"Comment,omitempty"`
	ReportedBy    string   `json:"reportedBy,omitempty"`
	VehicleNumber string   `json:"vehicleNumber,omitempty"`
}

type Report struct {
	ReportId       string
	ImageUrls      string
	Locations      string
	Offense        string
	SubmittedByRTO string
	Social         string
	TotalFine      int
	RTOApproved    string
	Comments       string
	ReportedBy     string
	VehicleNumber  string
}
