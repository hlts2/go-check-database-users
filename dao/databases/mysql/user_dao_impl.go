package mysql

import (
	"database/sql"

	"github.com/go-gorp/gorp"
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
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	defer dbmap.Db.Close()

	var users models.Users
	_, err = dbmap.Select(&users, "SELECT Host, User FROM mysql.user")
	if err != nil {
		return nil, err
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

	var user models.User
	err = stm.QueryRow(accountHost, accountName).Scan(&user.Host, &user.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &user, nil
}
