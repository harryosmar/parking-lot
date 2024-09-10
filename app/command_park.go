package app

import (
	"fmt"
	"regexp"
	"strings"
)

type Park struct {
	RegistrationNumber string
}

func (p *Park) Run(parkingLot *ParkingLot) {
	parkingLot.Park(p.RegistrationNumber)
}

// Generate : "park [car-number] [hours]"
func (p *Park) Generate(commandStr string) (Command, error) {
	commandStr = strings.ToLower(commandStr)
	re := regexp.MustCompile(`^[\s\t]*park[\s\t]+([a-z0-9\-]+)[\s\t]*$`)
	matches := re.FindAllStringSubmatch(commandStr, -1)

	if len(matches) == 1 && len(matches[0]) == 2 {
		registrationNumber := matches[0][1]
		return &Park{
			RegistrationNumber: strings.ToUpper(registrationNumber),
		}, nil
	}

	return nil, fmt.Errorf("command %s is not matched with 'Park Command'", commandStr)
}
