package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/hlts2/go-check-database-users/dao/databases/config"
	"github.com/hlts2/go-check-database-users/dao/factories"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List for database users",
	Run: func(cmd *cobra.Command, args []string) {
		if err := ls(cmd, args); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

//Execute ls command
func ls(cmd *cobra.Command, args []string) error {
	c := config.Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
	}

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()

	dao := factories.FactoryUserDao(dbms, c)
	if dao == nil {
		return errors.New("Invaild database management system")
	}

	users, err := dao.GetAllUsers()
	if err != nil {
		s.Stop()
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Host", "Name"})

	for _, v := range users {
		table.Append([]string{v.Host, v.Name})
	}

	s.Stop()
	table.Render()

	return nil
}
