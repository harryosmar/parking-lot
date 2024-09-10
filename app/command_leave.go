package app

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Leave struct {
	RegistrationNumber string
	DurationInHours    uint
}

func (l *Leave) Run(parkingLot *ParkingLot) {
	parkingLot.Leave(l.RegistrationNumber, l.DurationInHours)
}

// Generate : "leave [car-number] [hours]"
func (l *Leave) Generate(commandStr string) (Command, error) {
	commandStr = strings.ToLower(commandStr)
	re := regexp.MustCompile(`^[\s\t]*leave[\s\t]+([a-z0-9\-]+)[\s\t]+(\d+)[\s\t]*$`)
	matches := re.FindAllStringSubmatch(commandStr, -1)

	if len(matches) == 1 && len(matches[0]) == 3 {
		registrationNumber := matches[0][1]
		durationStr := matches[0][2]
		duration, err := strconv.ParseUint(durationStr, 10, 32)
		if err != nil {
			return nil, err
		}
		return &Leave{
			RegistrationNumber: strings.ToUpper(registrationNumber),
			DurationInHours:    uint(duration),
		}, nil
	}

	return nil, fmt.Errorf("command %s is not matched with 'Leave Command'", commandStr)
}

func (l *Leave) HelpUsage() string {
	return "leave [car-number] [hours] -> Removes (unpark) a car"
}
