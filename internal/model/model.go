package model

type User struct {
	Id       int
	Login    string
	Email    string
	Password string
}

type Post struct {
	Id       int
	AuthodID int
	Title    string
	BodyText string
}

type Comment struct {
	Id       int
	BodyText string
}
