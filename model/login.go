package model

// Login represents a login/password payload stored in records.
type Login struct {
	Login    string   `json:"login"`
	Password string   `json:"password"`
	Links    []string `json:"links"`
}
