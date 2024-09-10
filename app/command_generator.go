package app

import "strings"

type CommandGenerator struct {
	availableCommands []Command
}

func NewCommandGenerator(availableCommands []Command) *CommandGenerator {
	return &CommandGenerator{availableCommands: availableCommands}
}

func (c CommandGenerator) GenerateFromString(input string) []Command {
	commands := []Command{}

	// remove \r\n windows style
	inputCleaned := strings.ReplaceAll(input, "\r\n", "\n")
	lines := strings.Split(inputCleaned, "\n")

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		for _, command := range c.availableCommands {
			validCommand, err := command.Generate(line)
			if err == nil {
				commands = append(commands, validCommand)
				break
			}
		}
	}

	return commands
}
