package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Workflow struct {
	Name  string
	Rules []WorkflowRule
}

type WorkflowRule struct {
	Operand  byte
	Operator byte
	Value    int
	Action   string
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
	braceIndex := strings.Index(line, "{")
	name := line[0:braceIndex]
	workflowString := line[braceIndex+1 : len(line)-1]

	workflows := strings.Split(workflowString, ",")
	rules := make([]WorkflowRule, 0)
	for _, workflow := range workflows {

		colonIndex := strings.Index(workflow, ":")

		if colonIndex == -1 {
			rule := WorkflowRule{Operand: 0, Operator: 0, Value: 0, Action: workflow}
			rules = append(rules, rule)
			continue
		}

		operand := workflow[0]
		operator := workflow[1]

		numberString := workflow[2:colonIndex]
		value, _ := strconv.Atoi(numberString)
		action := workflow[colonIndex+1:]
		rule := WorkflowRule{Operand: operand, Operator: operator, Value: value, Action: action}
		rules = append(rules, rule)
	}

	return Workflow{Name: name, Rules: rules}

}

func main() {
	lines := ReadFile("test.txt")

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
		switch rule.Operator {
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

	return countAllCombinations(workflows, "in", map[byte][2]int{
		'x': {1, 4000},
		'm': {1, 4000},
		'a': {1, 4000},
		's': {1, 4000},
	})
}
func createMinMaxCopy(original map[byte][2]int) map[byte][2]int {
	newMinMax := make(map[byte][2]int, 4)
	for k, v := range original {
		newMinMax[k] = v
	}
	return newMinMax
}

func countAllCombinations(workflows map[string]Workflow, entry string, minmax map[byte][2]int) int {
	if entry == "R" {
		return 0
	} else if entry == "A" {
		result := 1
		for _, v := range minmax {
			result *= (v[1] - v[0] + 1)
		}
		return result
	}

	sum := 0
	for _, workflow := range workflows[entry].Rules {
		currentMinMax := minmax[workflow.Operand]

		// this behaves like a tree, so we
		// split the minmax part into 2 different ranges:
		// One that evaluates to true
		// One that evaluates to false
		// values differ for < and >
		// the new ranges to to +/- 1
		var truevalue [2]int
		var falsevalue [2]int

		if workflow.Operator == '<' {
			truevalue = [2]int{currentMinMax[0], workflow.Value - 1}
			falsevalue = [2]int{workflow.Value, currentMinMax[1]}

		} else if workflow.Operator == '>' {
			truevalue = [2]int{workflow.Value + 1, currentMinMax[1]}
			falsevalue = [2]int{currentMinMax[0], workflow.Value}
		} else {
			// if it's neither < nor > its just a direct jump to a new target, so calculate for that
			sum += countAllCombinations(workflows, workflow.Action, minmax)
			continue
		}

		if truevalue[0] <= truevalue[1] {
			// create a copy of current minmax values
			nextMinMax := createMinMaxCopy(minmax)
			// replace values for current operand with the new range that evaluates true
			nextMinMax[workflow.Operand] = truevalue
			// sum for next target rule
			sum += countAllCombinations(workflows, workflow.Action, nextMinMax)
		}

		// if low value is already higher than high, it's impossible, so break
		if falsevalue[0] > falsevalue[1] {
			break
		}

		// since we continued counting with the true value part, we need to continue with the false part
		minmax[workflow.Operand] = falsevalue

	}
	return sum
}

func ReadFile(fName string) []string {
	f, err := os.ReadFile(fName)
	if err != nil {
		panic("error reading file")
	}
	content := strings.Split(string(f), "\n")
	return content

}
