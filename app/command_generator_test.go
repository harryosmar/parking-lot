package app_test

import (
	"encoding/json"
	"github.com/harryosmar/parking-lot/app"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommandGenerateFromString(t *testing.T) {
	type args struct {
		input string
	}
	testData := []struct {
		name                string
		args                args
		expectedCommandsStr string
	}{
		{
			name: "TC1. Given empty multiple lines",
			args: args{
				input: `  	
			
    
		

`,
			},
			expectedCommandsStr: `[]`,
		},
		{
			name: "TC2. Given empty multiple lines with some valid lines",
			args: args{
				input: `  	
			


					    create 			     	6     
    status
		park    	KA-01-HH-1234   
		leave 			   DL-12-AA-9999     	2  		

`,
			},
			expectedCommandsStr: `[{"Size":6},{},{"RegistrationNumber":"KA-01-HH-1234"},{"RegistrationNumber":"DL-12-AA-9999","DurationInHours":2}]`,
		},
		{
			name: "TC3. Given multiple lines",
			args: args{
				input: `create 6
park KA-01-HH-1234
park KA-01-HH-9999
park KA-01-BB-0001
park KA-01-HH-7777
park KA-01-HH-2701
park KA-01-HH-3141
leave KA-01-HH-3141 4
status
park KA-01-P-333
park DL-12-AA-9999
leave KA-01-HH-1234 4
leave KA-01-BB-0001 6
leave DL-12-AA-9999 2
park KA-09-HH-0987
park CA-09-IO-1111
park KA-09-HH-0123
status`,
			},
			expectedCommandsStr: `[{"Size":6},{"RegistrationNumber":"KA-01-HH-1234"},{"RegistrationNumber":"KA-01-HH-9999"},{"RegistrationNumber":"KA-01-BB-0001"},{"RegistrationNumber":"KA-01-HH-7777"},{"RegistrationNumber":"KA-01-HH-2701"},{"RegistrationNumber":"KA-01-HH-3141"},{"RegistrationNumber":"KA-01-HH-3141","DurationInHours":4},{},{"RegistrationNumber":"KA-01-P-333"},{"RegistrationNumber":"DL-12-AA-9999"},{"RegistrationNumber":"KA-01-HH-1234","DurationInHours":4},{"RegistrationNumber":"KA-01-BB-0001","DurationInHours":6},{"RegistrationNumber":"DL-12-AA-9999","DurationInHours":2},{"RegistrationNumber":"KA-09-HH-0987"},{"RegistrationNumber":"CA-09-IO-1111"},{"RegistrationNumber":"KA-09-HH-0123"},{}]`,
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			commandGenerator := app.NewCommandGenerator([]app.Command{
				&app.Status{},
				&app.CreateSlot{},
				&app.Park{},
				&app.Leave{},
			})

			actualCommands := commandGenerator.GenerateFromString(tt.args.input)
			bytes, _ := json.Marshal(actualCommands)
			actualCommandsStr := string(bytes)

			assert.Equal(t, tt.expectedCommandsStr, actualCommandsStr, "The commands should be the same.")
		})
	}
}
