package domain

// User struct
type User struct {
	ID    uint64 `json:"id"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Email string `json:"email"`
}
