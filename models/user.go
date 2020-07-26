package models

type User struct {
	Id       int    `json:id`
	Username string `json:username`
	Password string `json:password`
}

type Creds struct {
	User string `json:user`
	Pass string `json:pass`
}
