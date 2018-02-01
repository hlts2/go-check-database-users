package mysql

import (
	"fmt"

	"github.com/hlts2/go-check-database-users/dao/databases/config"
)

type mysqlConfig struct {
	*config.Config
}

//NewMysqlConfig creates mysqlConfig instance
func NewMysqlConfig(c *config.Config) config.DBConfig {
	return &mysqlConfig{c}
}

//DSN returns database source name
func (c *mysqlConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/", c.User, c.Password, c.Host, c.Port)
}
