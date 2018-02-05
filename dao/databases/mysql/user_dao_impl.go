package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hlts2/go-check-database-users/dao/databases/config"
	"github.com/hlts2/go-check-database-users/models"
)

//UserDaoImpl is implementation of user dao interface
type UserDaoImpl struct {
	config.DBConfig
}

//GetAll returns User slice
func (u UserDaoImpl) GetAll() (models.Users, error) {
	db, err := sql.Open("tcp", u.DSN())
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT Host, User FROM mysql.user")
	if err != nil {
		return nil, err
	}

	users := make(models.Users, 1)
	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.Host, &user.Name)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
