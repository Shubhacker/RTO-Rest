package structs

type User struct {
	// Message   string `json:"authenticated"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserId    string `json:"userId"`
	JWT       string `json:"JWT"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginResponse struct {
	Authenticated bool   `json:"authenticated"`
	Message       string `json:"message"`
	UserInfo      User   `json:"user"`
}

// User          *UserInformation `json:"user,omitempty"`
// 	Authenticated *bool            `json:"authenticated,omitempty"`
// 	Message       *string          `json:"message,omitempty"`

type UserInfo struct {
	UserId    string
	UserName  string
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type UserRequest struct {
	UserName  string
	FirstName string
	LastName  string
	Email     string
	Password  string
}
