package cmd

import (
	"fmt"
	"os"

	"github.com/paralin/replayexport/pkg/csv"
	"github.com/spf13/cobra"
)

var csvParams struct {
	since     int
	minRating int
}

// csvCmd represents the csv command
var csvCmd = &cobra.Command{
	Use:   "csv",
	Short: "Export submissions to csv.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		sess, err := initMongo()
		defer sess.Close()
		if err != nil {
			fmt.Printf("Error connecting to mongo: %v\n", err)
			return
		}

		err = csv.ExportSubmissions(os.Stdout, sess.DB("replay").C("submissions"), csvParams.since, csvParams.minRating)
		if err != nil {
			fmt.Printf("Error exporting csv: %v\n", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(csvCmd)
	// csvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// csvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	csvCmd.Flags().IntVar(&csvParams.since, "since", 0, "unix timestamp for since (seconds)")
	csvCmd.Flags().IntVar(&csvParams.minRating, "min-rating", 0, "minimum rating")
}
