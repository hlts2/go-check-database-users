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
func (u UserDaoImpl) GetAllUsers() (models.Users, error) {
	db, err := sql.Open("mysql", u.DSN())
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

func (u UserDaoImpl) GetUser(accountName string, accountHost string) (*models.User, error) {
	db, err := sql.Open("mysql", u.DSN())
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stm, err := db.Prepare("SELECT Host, User From mysql.user WHERE Host = ? AND User = ?")
	if err != nil {
		return nil, err
	}
	defer stm.Close()

	user := new(models.User)
	stm.QueryRow(accountName, accountHost).Scan(user.Host, user.Name)

	return user, nil
}
