package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}
type Direction struct {
	X int
	Y int
}

type Workflow struct {
	Name  string
	Rules []WorkflowRule
}

type WorkflowRule struct {
	Operand   byte
	Operation byte
	Value     int
	Action    string
}

type Item struct {
	Values map[byte]int
}

func parseItem(line string) Item {
	line = line[1 : len(line)-1]
	fields := strings.Split(line, ",")
	item := Item{Values: make(map[byte]int, 0)}
	for _, field := range fields {
		number, _ := strconv.Atoi(field[2:])
		item.Values[field[0]] = number
	}
	return item
}

func parseWorkflow(line string) Workflow {
	//px{a<2006:qkq,m>2090:A,rfg}
	braceIndex := strings.Index(line, "{")
	name := line[0:braceIndex]
	workflowString := line[braceIndex+1 : len(line)-1]

	workflows := strings.Split(workflowString, ",")
	rules := make([]WorkflowRule, 0)
	for _, workflow := range workflows {

		colonIndex := strings.Index(workflow, ":")

		if colonIndex == -1 {
			rule := WorkflowRule{Operand: 0, Operation: 0, Value: 0, Action: workflow}
			rules = append(rules, rule)
			continue
		}

		operand := workflow[0]
		operator := workflow[1]

		numberString := workflow[2:colonIndex]
		value, _ := strconv.Atoi(numberString)
		action := workflow[colonIndex+1:]
		rule := WorkflowRule{Operand: operand, Operation: operator, Value: value, Action: action}
		rules = append(rules, rule)
	}

	return Workflow{Name: name, Rules: rules}

}

func main() {
	lines := ReadFile("test2.txt")

	fmt.Println("Puzzle1: ", puzzle1(lines))
	fmt.Println("Puzzle2: ", puzzle2(lines))
}

func parseAllLines(lines []string) (map[string]Workflow, []Item) {
	workflows := make(map[string]Workflow, 0)
	items := make([]Item, 0)
	parseItems := false
	for _, line := range lines {
		if line == "" {
			parseItems = true
			continue
		}
		if parseItems {
			items = append(items, parseItem(line))
		} else {
			workflow := parseWorkflow(line)
			workflows[workflow.Name] = workflow
		}

	}
	return workflows, items
}

func executeWorkflow(w Workflow, i Item) string {

	for _, rule := range w.Rules {
		switch rule.Operation {
		case '<':
			if i.Values[rule.Operand] < rule.Value {
				return rule.Action
			}

		case '>':
			if i.Values[rule.Operand] > rule.Value {
				return rule.Action
			}

		default:
			return rule.Action
		}
	}
	return "R"
}
func puzzle1(lines []string) int {
	workflows, items := parseAllLines(lines)

	//reduce workflow rules for all workflows
	// where all rules lead to the same destination
OUTER:
	for idx := range workflows {
		workflow := workflows[idx]
		action := workflow.Rules[0].Action
		for i := 1; i < len(workflow.Rules); i++ {
			if workflow.Rules[i].Action != action {
				continue OUTER
			}
			workflow.Rules = workflow.Rules[len(workflow.Rules)-1:]
			workflows[idx] = workflow
		}

	}

	accepted := make([]Item, 0)

	for _, item := range items {
		// always start with the workflow name "in"
		target := executeWorkflow(workflows["in"], item)

		for (target != "A") && (target != "R") {
			target = executeWorkflow(workflows[target], item)
		}

		if target == "A" {
			accepted = append(accepted, item)
		}

	}

	//fmt.Println("Rejected: ", len(rejected))

	//fmt.Println("Accepted: ", len(accepted))
	sumTotal := 0
	for _, item := range accepted {
		sumItem := 0
		for _, v := range item.Values {
			sumItem += v
		}
		sumTotal += sumItem
	}
	return sumTotal
}

func puzzle2(lines []string) int {
	workflows, _ := parseAllLines(lines)
	highest := make(map[byte]int)
	lowest := make(map[byte]int)

	highest['x'] = 4000
	highest['m'] = 4000
	highest['a'] = 4000
	highest['s'] = 4000

	for _, workflow := range workflows {
		for _, rule := range workflow.Rules {
			switch rule.Operation {
			case '<':
				val := highest[rule.Operand]
				if rule.Value > val {
					highest[rule.Operand] = rule.Value
				}
			case '>':
				val := lowest[rule.Operand]
				if rule.Value < val {
					lowest[rule.Operand] = rule.Value
				}
			}
		}
	}

	return 0
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ReadFile(fName string) []string {
	f, err := os.ReadFile(fName)
	if err != nil {
		panic("error reading file")
	}
	content := strings.Split(string(f), "\n")
	return content

}
