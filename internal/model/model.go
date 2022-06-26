package model

type User struct {
	Id        int    `json:"Id"`
	NickName  string `json:"NickName"`
	Age       int    `json:"Age"`
	Gender    string `json:"Gender"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
	Password  string `json:"Password"`
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
