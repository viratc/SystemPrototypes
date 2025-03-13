package models

type Group struct {
	Id          string
	ImageUrl    string
	Description string
	Users       []string
}

func NewGroup(id string, imageUrl string, description string, users []string) *Group {
	return &Group{Id: id, ImageUrl: imageUrl, Description: description, Users: users}
}

func (g *Group) GetUsers() []string {
	return g.Users
}
