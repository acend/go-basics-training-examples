package api

type User struct {
	Name      string `json:"name"`
	FullName  string `json:"full_name"`
	Followers int    `json:"followers"`
}
