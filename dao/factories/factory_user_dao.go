package factories

import (
	"github.com/hlts2/go-check-database-users/dao/databases/config"
	"github.com/hlts2/go-check-database-users/dao/databases/mysql"
	"github.com/hlts2/go-check-database-users/dao/interfaces"
)

//FactoryUserDao creates User Dao
func FactoryUserDao(s string, c *config.Config) interfaces.UserDao {
	var i interfaces.UserDao
	switch s {
	case "mysql":
		i = mysql.UserDaoImpl{
			mysql.NewMysqlConfig(c),
		}
	case "postgre":
		//TODO
	default:
		break
	}
	return i
}
