package app

import (
	"strings"
	"sync"
)

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

		command := c.resolveLineToCommand(line)
		if command != nil {
			commands = append(commands, command)
		}
	}

	return commands
}

func (c CommandGenerator) resolveLineToCommand(line string) Command {
	var wg sync.WaitGroup
	wg.Add(len(c.availableCommands))
	commandCh := make(chan Command)

	for _, command := range c.availableCommands {
		go func(command Command) {
			defer wg.Done()

			validCommand, err := command.Generate(line)
			if err == nil {
				commandCh <- validCommand
			}
		}(command)
	}

	go func() {
		wg.Wait()
		close(commandCh)
	}()

	for validCommand := range commandCh {
		return validCommand
	}

	return nil
}
