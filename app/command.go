package app

type Command interface {
	Run(parkingLot *ParkingLot)
	Generate(commandStr string) (Command, error)
	HelpUsage() string
}
