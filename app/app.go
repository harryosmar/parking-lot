package app

import "fmt"

type ParkingLot struct {
	firstNHours         uint
	costFirstNHours     float32
	costAdditionalHours float32
	spacesLen           uint
	tickets             map[uint]string
}

func NewParkingLot(firstNHours uint, costFirstNHours float32, costAdditionalHours float32) *ParkingLot {
	return &ParkingLot{
		firstNHours:         firstNHours,
		costFirstNHours:     costFirstNHours,
		costAdditionalHours: costAdditionalHours,
		spacesLen:           0,
		tickets:             map[uint]string{},
	}

}

func (p *ParkingLot) Run(commands ...Command) {
	for _, c := range commands {
		c.Run(p)
	}
}

func (p *ParkingLot) CreateSlot(size uint) {
	p.spacesLen = p.spacesLen + size
	fmt.Printf("Created parking lot with %d slots\n", p.spacesLen)
}

func (p *ParkingLot) Park(registrationNumber string) {
	if int(p.spacesLen) == len(p.tickets) {
		println("Sorry, parking lot is full")
		return
	}

	for _, existRegistrationNumber := range p.tickets {
		if existRegistrationNumber == registrationNumber {
			fmt.Printf("Registration Number %s already registered\n", registrationNumber)
			return
		}
	}

	var i uint
	for i = 1; i <= p.spacesLen; i++ {
		if _, found := p.tickets[i]; !found {
			p.tickets[i] = registrationNumber
			fmt.Printf("Allocated slot number: %d\n", i)
			return
		}
	}
}

func (p *ParkingLot) Leave(inputRegistrationNumber string, durationInHour uint) {
	for slotNumber, registrationNumber := range p.tickets {
		if registrationNumber == inputRegistrationNumber {
			cost := p.CalculateCost(durationInHour)
			delete(p.tickets, slotNumber)

			fmt.Printf("Registration Number %s from Slot %d has left with Charge %.0f\n", registrationNumber, slotNumber, cost)
			return
		}
	}

	fmt.Printf("Registration Number %s not found\n", inputRegistrationNumber)
}

func (p *ParkingLot) CalculateCost(durationInHour uint) float32 {
	charge := float32(durationInHour) * p.costFirstNHours
	if durationInHour > p.firstNHours {
		charge = float32(p.firstNHours)*p.costFirstNHours + (float32(durationInHour-p.firstNHours) * p.costAdditionalHours)
	}

	return charge
}

func (p *ParkingLot) Status() string {
	result := "Slot No. Registration No.\n"

	var i uint
	for i = 1; i <= p.spacesLen; i++ {
		slotStr := fmt.Sprintf("%d	-\n", i)
		if registrationNumber, found := p.tickets[i]; found {
			slotStr = fmt.Sprintf("%d	%s\n", i, registrationNumber)
		}

		result = fmt.Sprintf("%s%s", result, slotStr)
	}

	print(result)

	return result
}
