package app

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type CreateSlot struct {
	Size uint
}

func (c *CreateSlot) Run(parkingLot *ParkingLot) {
	parkingLot.CreateSlot(c.Size)
}

// Generate : "create [size]"
func (c *CreateSlot) Generate(commandStr string) (Command, error) {
	commandStr = strings.ToLower(commandStr)
	re := regexp.MustCompile(`^[\s\t]*create[\s\t]+(\d+)[\s\t]*$`)
	matches := re.FindAllStringSubmatch(commandStr, -1)

	if len(matches) == 1 && len(matches[0]) == 2 {
		sizeStr := matches[0][1]
		size, err := strconv.ParseUint(sizeStr, 10, 32)
		if err != nil {
			return nil, err
		}
		return &CreateSlot{Size: uint(size)}, nil
	}

	return nil, fmt.Errorf("command %s is not matched with 'CreateSlot Command'", commandStr)
}

func (c *CreateSlot) HelpUsage() string {
	return "create [size] - Creates parking lot of size n"
}
