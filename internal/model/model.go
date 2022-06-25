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
	Tags     []Tag
	Comments []Comment
}

type Comment struct {
	Id       int
	AuthodID int
	PostId   int
	BodyText string
}

type Tag struct {
	Id     int
	PostId int
	Value  string
}
