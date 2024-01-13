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
	debug = false

	MAX_INT = 2147483647

	MIN_VAL = 1
	MAX_VAL = 4000
)

type Part struct {
	x, m, a, s int
}

func (p *Part) String() string {
	return fmt.Sprintf("{x=%v, m=%v, a=%v, s=%v}", p.x, p.m, p.a, p.s)
}

type Rule struct {
	left string
	comp string
	right int
	action string
}

func (r *Rule) Apply(p *Part) string {
	if debug {
		fmt.Printf("checking rule %v against part %v\n", r, p)
	}
	if r.left == "" {
		if debug {
			fmt.Printf("ending action: %v\n", r.action)
		}
		return r.action
	}

	var left int
	switch r.left {
	case "x":
		left = p.x
	case "m":
		left = p.m
	case "a":
		left = p.a
	case "s":
		left = p.s
	}

	switch r.comp {
	case "<":
		if left < r.right {
			if debug {
				fmt.Printf("left < right: [%v]\n", r.action)
			}
			return r.action
		} else {
			if debug {
				fmt.Printf("left < right: []\n")
			}
			return ""
		}
	case ">":
		if left > r.right {
			if debug {
				fmt.Printf("left > right: [%v]\n", r.action)
			}
			return r.action
		} else {
			if debug {
				fmt.Printf("left > right: []\n")
			}
			return ""
		}
	}

	if debug {
		fmt.Printf("should not reach here: []\n")
	}
	return ""
}

func (r *Rule) String() string {
	if r.left == "" {
		return fmt.Sprintf("{-> %v}", r.action)
	}
	return fmt.Sprintf("{%v %v %v -> %v}", r.left, r.comp, r.right, r.action)
}

var workflows map[string][]*Rule
var parts []*Part

func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	workflows = make(map[string][]*Rule)

	parsingParts := false
	for _, line := range inputAsLines {
		if line == "" {
			parsingParts = true
			continue
		}

		if !parsingParts {
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
		} else {
			rawParts := strings.ReplaceAll(strings.ReplaceAll(line, "{", ""), "}", "")

			part := &Part{}
			for _, rawPart := range strings.Split(rawParts, ",") {
				s := strings.Split(rawPart, "=")
				r, _ := strconv.Atoi(s[1])

				switch s[0] {
				case "x":
					part.x = r
				case "m":
					part.m = r
				case "a":
					part.a = r
				case "s":
					part.s = r
				}
			}

			parts = append(parts, part)

		}
	}

	fmt.Printf("workflows = %v\n", workflows)
	fmt.Printf("parts = %v\n", parts)

	res := int64(0)
	for _, p := range parts {
		a := processPart(p)
		fmt.Printf("%v -> %v\n", p, a)
		if a == "A" {
			res += int64(p.x) + int64(p.m) + int64(p.a) + int64(p.s)
		}
	}
	fmt.Printf("res = %v\n", res)
}

func processPart(p *Part) string {
	workflow := workflows["in"]
	for {		
		if debug {
			fmt.Printf("checking workflow: %v\n", workflow)
		}
		var res string
		for _, r := range workflow {
			res = r.Apply(p)

			switch res {
			case "A":
				return "A"
			case "R":
				return "R"
			}

			if len(res) > 0 {
				break
			}

		}

		workflow = workflows[res]
	}

	return ""
}