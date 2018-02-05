package interfaces

import (
	"github.com/hlts2/go-check-database-users/models"
)

//UserDao is user table dao interface
type UserDao interface {
	GetAllUsers() (models.Users, error)
	GetUser(*models.User, error)
}
