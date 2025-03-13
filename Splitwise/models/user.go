package models

type User struct {
	Id        string
	FirstName string
	LastName  string
	Bio       string
	ImageUrl  string
}

func NewUser(id string, firstName string, lastName string, bio string, imageUrl string) *User {
	return &User{Id: id, FirstName: firstName, LastName: lastName, Bio: bio, ImageUrl: imageUrl}
}
