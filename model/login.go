package model

type Login struct {
	Login    string   `json:"login"`
	Password string   `json:"password"`
	Links    []string `json:"links"`
}
