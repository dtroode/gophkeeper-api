package model

type URL string

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Links    []URL  `json:"links"`
}
