package cmd

import (
	"fmt"
	"os"

	"github.com/hlts2/go-check-database-users/dao/databases/config"
	"github.com/hlts2/go-check-database-users/dao/factories"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List for Mysql users",
	Run: func(cmd *cobra.Command, args []string) {
		if err := ls(cmd, args); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

var host string
var port int
var user string
var password string

func init() {
	lsCmd.Flags().StringVarP(&host, "Host", "H", "localhost", "Hostname（localhost）")
	lsCmd.Flags().IntVarP(&port, "port", "p", 3306, "Port（3306）")
	lsCmd.Flags().StringVarP(&user, "user", "u", "root", "Username（root)")
	lsCmd.Flags().StringVarP(&password, "password", "P", "", "Password")
	rootCmd.AddCommand(lsCmd)
}

func ls(cmd *cobra.Command, args []string) error {
	c := config.Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
	}

	dao := factories.FactoryUserDao("mysql", &c)

	users, err := dao.GetAll()
	if err != nil {
		return err
	}

	fmt.Println(users)
	return nil
}
