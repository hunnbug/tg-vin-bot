package models

type User struct {
	firstName string
	lastName  string
	userName  string
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) UserName() string {
	return u.userName
}

func CreateUser(firstName string, lastName string, userName string) *User {
	return &User{
		firstName,
		lastName,
		userName,
	}
}
