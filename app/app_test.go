package app_test

import (
	"github.com/harryosmar/parking-lot/app"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunCommands(t *testing.T) {
	type args struct {
		commands   []app.Command
		parkingLot *app.ParkingLot
	}
	testData := []struct {
		name           string
		args           args
		expectedStatus string
	}{
		{
			name: "TC1. Given multiple commands When call parkingLot.Run Then status should be valid",
			args: args{
				commands: []app.Command{
					&app.CreateSlot{Size: 6},
					&app.Park{RegistrationNumber: "KA-01-HH-1234"},
					&app.Park{RegistrationNumber: "KA-01-HH-1234"},
					&app.Park{RegistrationNumber: "KA-01-HH-9999"},
					&app.Park{RegistrationNumber: "KA-01-BB-0001"},
					&app.Park{RegistrationNumber: "KA-01-HH-7777"},
					&app.Park{RegistrationNumber: "KA-01-HH-2701"},
					&app.Park{RegistrationNumber: "KA-01-HH-3141"},
					&app.Leave{RegistrationNumber: "KA-01-HH-3141", DurationInHours: 4},
					&app.Status{},
					&app.Park{RegistrationNumber: "KA-01-P-333"},
					&app.Park{RegistrationNumber: "DL-12-AA-9999"},
					&app.Leave{RegistrationNumber: "KA-01-HH-1234", DurationInHours: 4},
					&app.Leave{RegistrationNumber: "KA-01-BB-0001", DurationInHours: 6},
					&app.Leave{RegistrationNumber: "DL-12-AA-9999", DurationInHours: 2},
					&app.Park{RegistrationNumber: "KA-09-HH-0987"},
					&app.Park{RegistrationNumber: "CA-09-IO-1111"},
					&app.Park{RegistrationNumber: "KA-09-HH-0123"},
					&app.Status{},
				},
				parkingLot: app.NewParkingLot(2, 10, 5),
			},
			expectedStatus: `Slot No.	Registration No.
1		KA-09-HH-0987
2		KA-01-HH-9999
3		CA-09-IO-1111
4		KA-01-HH-7777
5		KA-01-HH-2701
6		KA-01-P-333
`,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.parkingLot.Run(tt.args.commands...)
			actualStatus := tt.args.parkingLot.Status()
			assert.Equal(t, tt.expectedStatus, actualStatus, "The two status should be the same.")
		})
	}
}

func TestCalculateCost(t *testing.T) {
	type args struct {
		parkingLot     *app.ParkingLot
		durationInHour uint
	}
	testData := []struct {
		name         string
		args         args
		expectedCost float32
	}{
		{
			name: "TC1. Given duration 1 Hour When call CalculateCost Then return 10",
			args: args{
				parkingLot:     app.NewParkingLot(2, 10, 5),
				durationInHour: 1,
			},
			expectedCost: 10,
		},
		{
			name: "TC2. Given duration 2 Hour When call CalculateCost Then return 20",
			args: args{
				parkingLot:     app.NewParkingLot(2, 10, 5),
				durationInHour: 2,
			},
			expectedCost: 20,
		},
		{
			name: "TC3. Given duration 3 Hour When call CalculateCost Then return 25",
			args: args{
				parkingLot:     app.NewParkingLot(2, 10, 5),
				durationInHour: 3,
			},
			expectedCost: 25,
		},
		{
			name: "TC4. Given duration 4 Hour When call CalculateCost Then return 30",
			args: args{
				parkingLot:     app.NewParkingLot(2, 10, 5),
				durationInHour: 4,
			},
			expectedCost: 30,
		},
		{
			name: "TC5. Given duration 10 Hour When call CalculateCost Then return 60",
			args: args{
				parkingLot:     app.NewParkingLot(2, 10, 5),
				durationInHour: 10,
			},
			expectedCost: 60,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			actualCost := tt.args.parkingLot.CalculateCost(tt.args.durationInHour)
			assert.Equal(t, tt.expectedCost, actualCost, "The cost should be the same.")
		})
	}
}
