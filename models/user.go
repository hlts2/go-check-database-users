package models

//User represents mysq.user table
type User struct {
	Host string
	Name string
}

//Users is User slice
type Users []User
