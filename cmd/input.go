package cmd

import (
	"fmt"
	"github.com/harryosmar/parking-lot/app"
	"github.com/harryosmar/parking-lot/config"
	"os"

	"github.com/spf13/cobra"
)

// inputCmd represents the input command
var inputCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "input /path/to/file.txt",
	Short: "Consume input from file contains commands",
	Long: `The system accept a filename as a parameter and read the commands from that file.
List of Commands :
create [size] - Creates parking lot of size n
park [car-number] - Parks a car
leave [car-number] [hours] -> Removes (unpark) a car
status -> Prints status of the parking lot

Examples of Input (content of file) :
create 6
park KA-01-HH-1234
leave KA-01-HH-3141 4
status
park KA-09-HH-0987
`,
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		err := processFile(filePath)
		if err != nil {
			fmt.Printf("Error processing file: %v\n", err)
			return
		}
	},
}

var (
	parkingLot       *app.ParkingLot
	commandGenerator *app.CommandGenerator
)

func init() {
	cfg := config.Get()
	parkingLot = app.NewParkingLot(cfg.FirstNHour, cfg.CostFirstNHour, cfg.CostAdditionalHour)
	commandGenerator = app.NewCommandGenerator([]app.Command{
		&app.Status{},
		&app.CreateSlot{},
		&app.Park{},
		&app.Leave{},
	})
	rootCmd.AddCommand(inputCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inputCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inputCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func processFile(filePath string) error {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	parkingLot.Run(
		commandGenerator.GenerateFromString(
			string(bytes),
		)...,
	)

	return nil
}
