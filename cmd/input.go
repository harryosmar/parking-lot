package cmd

import (
	"fmt"
	"github.com/harryosmar/parking-lot/app"
	"github.com/harryosmar/parking-lot/config"
	"github.com/spf13/cobra"
	"os"
)

// inputCmd represents the input command
var inputCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "input /path/to/file.txt",
	Short: "Consume input from file contains commands",
	Long: fmt.Sprintf(
		"The system accept a filename as a parameter and read the commands from that file.\nList of Commands :\n%s\n%s\n%s\n%s",
		(&app.CreateSlot{}).HelpUsage(),
		(&app.Park{}).HelpUsage(),
		(&app.Leave{}).HelpUsage(),
		(&app.Status{}).HelpUsage(),
	),
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
