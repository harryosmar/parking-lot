package app

import (
	"fmt"
	"regexp"
	"strings"
)

type Status struct {
}

func (s *Status) Run(parkingLot *ParkingLot) {
	parkingLot.Status()
}

// Generate : "status"
func (s *Status) Generate(commandStr string) (Command, error) {
	commandStr = strings.ToLower(commandStr)
	re := regexp.MustCompile(`^[\s\t]*status[\s\t]*$`)
	matches := re.FindAllStringSubmatch(commandStr, -1)

	if len(matches) == 1 && len(matches[0]) == 1 {
		return &Status{}, nil
	}

	return nil, fmt.Errorf("command %s is not matched with 'Park Command'", commandStr)
}
