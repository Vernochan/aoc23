package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	FlipFlop    = '%'
	Conjunction = '&'
	High        = true
	Low         = false
)

type Module struct {
	Name    string
	Type    rune
	Inputs  []string
	Outputs []string
	Memory  map[string]bool
}

type Pulse struct {
	Source  string
	Value   bool
	Targets []string
}

func main() {
	lines := ReadFile("test2.txt")

	fmt.Println("Puzzle1: ", puzzle1(lines))
	fmt.Println("Puzzle2: ", puzzle2(lines))
}

func parseAllLines(lines []string) map[string]Module {

	modules := make(map[string]Module, 0)

	for _, line := range lines {
		newModule := Module{Outputs: make([]string, 0), Inputs: make([]string, 0)}
		idx := strings.Index(line, " ")

		name := line[0:idx]

		if name[0] == FlipFlop || name[0] == Conjunction {
			newModule.Name = line[1:idx]
			newModule.Type = rune(name[0])
		} else {
			newModule.Name = name
		}

		idx = strings.Index(line, "->") + 3

		targetString := line[idx:]

		targets := strings.Split(targetString, ", ")

		newModule.Outputs = append(newModule.Outputs, targets...)

		modules[newModule.Name] = newModule
	}

	for _, module := range modules {
		for _, target := range module.Outputs {
			if _, ok := modules[target]; !ok {
				// if output module does not exists
				continue
			}
			tmp := modules[target]
			tmp.Inputs = append(tmp.Inputs, module.Name)
			modules[target] = tmp
		}
	}
	return modules
}

func puzzle1(lines []string) int {
	modules := parseAllLines(lines)
	state := make(map[string]bool)

	processingSignals := make([]Pulse, 0)
	processingSignals = append(processingSignals, Pulse{Value: Low, Targets: []string{"broadcaster"}})

	for len(processingSignals) > 0 {
		nextSignals := make([]Pulse, 0)
		for _, signal := range processingSignals {
			nextSignals = append(nextSignals, processSignal(modules, state, signal)...)
		}

		processingSignals = nextSignals
	}

	return len(modules)
}

func puzzle2(lines []string) int {

	//workflows, _ := parseAllLines(lines)

	return 0
}

func processSignal(modules map[string]Module, state map[string]bool, signal Pulse) []Pulse {
	output := make([]Pulse, 0)
	for _, target := range signal.Targets {

		if _, ok := modules[target]; !ok {
			// if output module does not exists
			continue
		}
		module := modules[target]

		switch module.Type {
		case Conjunction:
			newSignal := Pulse{Source: module.Name, Targets: module.Outputs, Value: Low}
			module.Memory[signal.Source] = signal.Value
			for _, v := range module.Inputs {
				if !module.Memory[v] {
					newSignal.Value = High
				}
			}

		case FlipFlop:

		default:
			output = append(output, Pulse{Value: signal.Value, Targets: module.Outputs, Source: module.Name})

		}

	}
	return output
}

func ReadFile(fName string) []string {
	f, err := os.ReadFile(fName)
	if err != nil {
		panic("error reading file")
	}
	content := strings.Split(string(f), "\n")
	return content

}
