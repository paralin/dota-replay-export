package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/mgo.v2"
)

var mongoUrl string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "replayexport",
	Short: "Exports stuff from DOTAReplay to csv.",
	Long:  ``,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&mongoUrl, "mongo", "", "mongo server")
}

func initMongo() (*mgo.Session, error) {
	return mgo.Dial(mongoUrl)
}
