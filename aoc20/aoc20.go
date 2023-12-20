package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	FlipFlop    = '%'
	Conjunction = '&'
	RX          = '>'
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
	lines := ReadFile("test.txt")

	fmt.Println("Puzzle1: ", puzzle1(lines))
	fmt.Println("Puzzle2: ", puzzle2(lines))
}

func parseAllLines(lines []string) map[string]Module {

	modules := make(map[string]Module, 0)

	for _, line := range lines {
		newModule := Module{Outputs: make([]string, 0), Inputs: make([]string, 0), Memory: make(map[string]bool)}
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
	count := make(map[bool]int, 2)
	processingSignals := make([]Pulse, 0)

	for i := 0; i < 1000; i++ {
		processingSignals = append(processingSignals, Pulse{Value: Low, Targets: []string{"broadcaster"}})

		for len(processingSignals) > 0 {
			nextSignals := make([]Pulse, 0)
			for _, signal := range processingSignals {
				count[signal.Value] += len(signal.Targets)
				nextSignals = append(nextSignals, processSignal(modules, state, signal)...)
			}

			processingSignals = nextSignals
		}

	}
	retval := 1
	for _, v := range count {
		retval *= v
	}

	return retval
}

func puzzle2(lines []string) int {

	modules := parseAllLines(lines)

	rxSource := ""

	// search node that terminates in "rx"
	// because only a single node writes to "rx"
	for _, module := range modules {
		for _, target := range module.Outputs {
			if target == "rx" {
				rxSource = module.Name
			}
		}
	}

	destTargets := make([]string, 0)

	// the node that terminates in "rx" is a conjunction,
	// so we need all modules, that terminate there
	for _, module := range modules {
		for _, target := range module.Outputs {
			if target == rxSource {
				destTargets = append(destTargets, module.Name)
			}
		}
	}

	cycles := make(map[string]int, 4)

	state := make(map[string]bool)
	processingSignals := make([]Pulse, 0)

	// only run as often as is needed to find the cycle length for all destTargets
	for i := 1; len(cycles) < len(destTargets); i++ {

		processingSignals = append(processingSignals, Pulse{Value: Low, Targets: []string{"broadcaster"}})

		for len(processingSignals) > 0 {
			nextSignals := make([]Pulse, 0)
			for _, signal := range processingSignals {
				nextSignals = append(nextSignals, processSignal(modules, state, signal)...)

				for _, input := range destTargets {

					// if current target doesn't have an entry in cycles (aka first time it happened) directly inspect memory of rxSource
					if _, ok := cycles[input]; !ok && modules[rxSource].Memory[input] {
						cycles[input] = i
					}

				}

			}

			processingSignals = nextSignals
		}

	}
	nums := make([]int, 0, len(cycles))
	for _, c := range cycles {
		nums = append(nums, c)
	}
	return leastCommonMultiple(nums...)
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func leastCommonMultiple(integers ...int) int {
	if len(integers) < 2 {
		panic("not enough numbers for LCM")
	}
	result := integers[0] * integers[1] / greatestCommonDivisor(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = leastCommonMultiple(result, integers[i])
	}

	return result
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
					break
				}
			}
			output = append(output, newSignal)

		case FlipFlop:
			if !signal.Value {
				module.Memory["status"] = !module.Memory["status"]
				newSignal := Pulse{Source: module.Name, Targets: module.Outputs, Value: module.Memory["status"]}
				output = append(output, newSignal)
			}

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
