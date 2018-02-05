package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/hlts2/go-check-database-users/dao/databases/config"
	"github.com/hlts2/go-check-database-users/dao/factories"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-check-mysql-users",
	Short: "A CLI tool for checking mysql user",
	Run: func(cmd *cobra.Command, args []string) {
		root(cmd, args)
	},
}

var host string
var port int
var user string
var password string
var accountName string
var accountHost string

func init() {
	rootCmd.PersistentFlags().StringVarP(&host, "Host", "H", "localhost", "Hostname（localhost）")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 3306, "Port（3306）")
	rootCmd.PersistentFlags().StringVarP(&user, "user", "u", "root", "Username（root)")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "P", "", "Password")
	rootCmd.Flags().StringVarP(&accountName, "account-name", "a", "", "account user name")
	rootCmd.Flags().StringVarP(&accountHost, "account-Host", "n", "", "account host name")
}

//Execute root command
func root(cmd *cobra.Command, args []string) {
	c := config.Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
	}

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	defer s.Stop()

	dao := factories.FactoryUserDao("mysql", &c)
	ok := dao.IsConnect()
	if ok {
		//TODO write message
	}
}

//Execute run command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
