package model

type User struct {
	Id       int
	Login    string
	Email    string
	Password string
}

type Posts struct {
	Id       int
	Title    string
	BodyText string
}

type Comments struct {
	Id       int
	BodyText string
}
