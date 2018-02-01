package interfaces

import (
	"github.com/hlts2/go-check-database-users/models"
)

//UserDao is User Dao interface
type UserDao interface {
	GetAll() (models.Users, error)
}
