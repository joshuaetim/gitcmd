package cli

import "time"

type Issue struct {
	Number		int
	Title		string
	Body		string
	HTMLURL		string 		`json:"html_url"`
	User		*User
	State		string
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

type User struct {
	Login 	string
}

const Url = "https://api.github.com/repos"