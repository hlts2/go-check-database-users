package models

//User represents mysq.user table
type User struct {
	Host string `db:"Host"`
	Name string `db:"User"`
}

//Users is User slice
type Users []User
