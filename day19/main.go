package main

import (
    "fmt"
    _ "embed"
    "strings"
    "strconv"
)

//go:embed input
var input string

const (
	debug = true

	MAX_INT = 2147483647

	MIN_VAL = 1
	MAX_VAL = 4000
)

type Range struct {
	p, q int
}

func (r Range) String() string {
	return fmt.Sprintf("[%v, %v]", r.p, r.q)
}

type Part struct {
	x, m, a, s Range
}

func (p *Part) String() string {
	return fmt.Sprintf("{x=%v, m=%v, a=%v, s=%v}", p.x, p.m, p.a, p.s)
}

func (p *Part) Copy() *Part {
	return &Part{
		x: p.x, m: p.m, a: p.a, s: p.s,
	}
}

type Rule struct {
	left string
	comp string
	right int
	action string
}

func (r *Rule) String() string {
	if r.left == "" {
		return fmt.Sprintf("{-> %v}", r.action)
	}
	return fmt.Sprintf("{%v %v %v -> %v}", r.left, r.comp, r.right, r.action)
}

var workflows map[string][]*Rule

func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	workflows = make(map[string][]*Rule)

	for _, line := range inputAsLines {
		if line == "" {
			break
		}

		name := line[:strings.Index(line, "{")]
		rawWorkflows := strings.ReplaceAll(strings.ReplaceAll(line, name + "{", ""), "}", "")
		for _, rawWorkflow := range strings.Split(rawWorkflows, ",") {
			if strings.Index(rawWorkflow, ":") < 0 {
				workflows[name] = append(workflows[name], &Rule{left: "", comp: "", right: 0, action: rawWorkflow})
				continue
			}

			s := strings.Split(rawWorkflow, ":")
			r, _ := strconv.Atoi(s[0][2:])
			workflows[name] = append(workflows[name], &Rule{left: string(s[0][0]), comp: string(s[0][1]), right: r, action: s[1]})
		}
	}

	fmt.Printf("workflows = %v\n", workflows)

	acceptedParts := processPart(&Part{
		x: Range{p: MIN_VAL, q: MAX_VAL},
		m: Range{p: MIN_VAL, q: MAX_VAL},
		a: Range{p: MIN_VAL, q: MAX_VAL},
		s: Range{p: MIN_VAL, q: MAX_VAL},
	}, "in", 0)

	fmt.Printf("acceptedParts = %v\n", acceptedParts)

	res := int64(0)
	for _, p := range acceptedParts {
		res += int64(p.x.q - p.x.p + 1) * int64(p.m.q - p.m.p + 1) * int64(p.a.q - p.a.p + 1) * int64(p.s.q - p.s.p + 1)
	}
	fmt.Printf("res = %v\n", res)
}

func processPart(p *Part, workflowName string, step int) []*Part {
	// if debug {
	// 	fmt.Printf("processPart(%v, %v, %v)\n", p, workflowName, step)
	// }
	workflow := workflows[workflowName]
	r := workflow[step]
	if debug {
		fmt.Printf("part: %v, workflow(%v): %v, rule: %v\n", p, workflowName, workflow, r)
	}

	if r.left == "" {
		switch r.action {
		case "R":
			return []*Part{}
		case "A":
			return []*Part{p}
		default:
			return processPart(p, r.action, 0)
		}
	}

	if r.comp == ">" {
		matchingPart := p.Copy()
		missedPart := p.Copy()
		switch r.left {
		case "x":
			missedPart.x = Range{p: p.x.p, q: r.right}
			matchingPart.x = Range{p: r.right + 1, q: p.x.q}
		case "m":
			missedPart.m = Range{p: p.m.p, q: r.right}
			matchingPart.m = Range{p: r.right + 1, q: p.m.q}
		case "a":
			missedPart.a = Range{p: p.a.p, q: r.right}
			matchingPart.a = Range{p: r.right + 1, q: p.a.q}
		case "s":
			missedPart.s = Range{p: p.s.p, q: r.right}
			matchingPart.s = Range{p: r.right + 1, q: p.s.q}
		}

		var res []*Part
		res = append(res, processPart(missedPart, workflowName, step + 1)...)

		switch r.action {
		case "R":
		case "A":
			res = append(res, matchingPart)
		default:
			res = append(res, processPart(matchingPart, r.action, 0)...)
		}
		return res
	} else { // r.comp == "<"
		matchingPart := p.Copy()
		missedPart := p.Copy()
		switch r.left {
		case "x":
			matchingPart.x = Range{p: p.x.p, q: r.right - 1}
			missedPart.x = Range{p: r.right, q: p.x.q}
		case "m":
			matchingPart.m = Range{p: p.m.p, q: r.right - 1}
			missedPart.m = Range{p: r.right, q: p.m.q}
		case "a":
			matchingPart.a = Range{p: p.a.p, q: r.right - 1}
			missedPart.a = Range{p: r.right, q: p.a.q}
		case "s":
			matchingPart.s = Range{p: p.s.p, q: r.right - 1}
			missedPart.s = Range{p: r.right, q: p.s.q}
		}

		var res []*Part
		res = append(res, processPart(missedPart, workflowName, step + 1)...)

		switch r.action {
		case "R":
		case "A":
			res = append(res, matchingPart)
		default:
			res = append(res, processPart(matchingPart, r.action, 0)...)
		}
		return res
	}
}