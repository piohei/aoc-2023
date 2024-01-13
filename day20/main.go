package main

import (
    "fmt"
    _ "embed"
    "strings"
)

//go:embed input
var input string

const (
	debug = false
	buttonPushedTimes = 1000

	MAX_INT = 2147483647

	MIN_VAL = 1
	MAX_VAL = 4000
)

type Pulse int
const (
	LOW = iota
	HIGH = iota
)

func (p Pulse) String() string {
	if p == LOW {
		return "low"
	} else if p == HIGH {
		return "high"
	} else {
		return "unknown"
	}
}

type Event struct {
	from string
	to string
	pulse Pulse
}

type Module interface {
	Process(from string, p Pulse) []Event
	AddOutput(name string)
	AddInput(name string)
	IsInitial() bool
}

type CommonModule struct {
	name string
	input []string
	output []string
}

func (cm *CommonModule) AddOutput(name string) {
	cm.output = append(cm.output, name)
}

func (cm *CommonModule) AddInput(name string) {
	cm.input = append(cm.input, name)
}

func (cm *CommonModule) createEvents(p Pulse) []Event {
	var res []Event

	for _, n := range cm.output {
		res = append(res, Event{from: cm.name, to: n, pulse: p})
	}

	return res
}

func (cm *CommonModule) String() string {
	return fmt.Sprintf("Module(name=%v, input=%v, output=%v)", cm.name, cm.input, cm.output)
}

type FlipFlop struct {
	*CommonModule
	on bool
}

func NewFlipFlop(name string) *FlipFlop {
	return &FlipFlop{
		CommonModule: &CommonModule{
			name: name,
		},
		on: false,
	}
}

func (ff *FlipFlop) Process(from string, p Pulse) []Event {
	var res []Event
	if p == HIGH {
		return res
	}

	if ff.on {
		ff.on = false
		return ff.createEvents(LOW)
	} else {
		ff.on = true
		return ff.createEvents(HIGH)
	}
}

func (ff *FlipFlop) IsInitial() bool {
	return ff.on == false
}

type Conjuction struct {
	*CommonModule
	memory map[string]Pulse
}

func NewConjuction(name string) *Conjuction {
	return &Conjuction{
		CommonModule: &CommonModule{
			name: name,
		},
		memory: make(map[string]Pulse),
	}
}

func (c *Conjuction) AddInput(name string) {
	c.CommonModule.AddInput(name)
	c.memory[name] = LOW
}

func (c *Conjuction) Process(from string, p Pulse) []Event {
	c.memory[from] = p

	for _, v := range c.memory {
		if v == LOW {
			return c.createEvents(HIGH)
		}
	}

	return c.createEvents(LOW)
}

func (c *Conjuction) IsInitial() bool {
	for _, v := range c.memory {
		if v == HIGH {
			return false
		}
	}
	return true
}

type Broadcast struct {
	*CommonModule
}

func NewBroadcast(name string) *Broadcast {
	return &Broadcast{
		CommonModule: &CommonModule{
			name: name,
		},
	}
}

func (b *Broadcast) Process(from string, p Pulse) []Event {
	return b.createEvents(p)
}

func (b *Broadcast) IsInitial() bool {
	return true
}

type Button struct {
	*CommonModule
}

func NewButton(name string) *Button {
	return &Button{
		CommonModule: &CommonModule{
			name: name,
			output: []string{"broadcaster"},
		},
	}
}

func (b *Button) Process(from string, p Pulse) []Event {
	return b.createEvents(LOW)
}

func (b *Button) IsInitial() bool {
	return true
}

type RX struct {
	*CommonModule
	on bool
}

func NewRX(name string) *RX {
	return &RX{
		CommonModule: &CommonModule{
			name: name,
			output: []string{},
		},
		on: false,
	}
}

func (rx *RX) Process(from string, p Pulse) []Event {
	if p == LOW {
		rx.on = true
	}
	return []Event{}
}

func (rx *RX) IsInitial() bool {
	return true
}

func (rx *RX) IsOn() bool {
	return rx.on
}

var modules map[string]Module
var lookingForEvent Event
var eventOccurred bool

func main() {
	fmt.Println(input)
	inputAsLines := strings.Split(input, "\n")

	modules = make(map[string]Module)
	inputs := make(map[string][]string)

	for _, line := range inputAsLines {
		s := strings.Split(line, " -> ")

		switch {
		case string(s[0][0]) == "%":
			name := strings.ReplaceAll(s[0], "%", "")

			modules[name] = NewFlipFlop(name)

			for _, out := range strings.Split(s[1], ",") {
				out = strings.ReplaceAll(out, " ", "")
				modules[name].AddOutput(out)
				inputs[out] = append(inputs[out], name)
			}
		case string(s[0][0]) == "&":
			name := strings.ReplaceAll(s[0], "&", "")

			modules[name] = NewConjuction(name)

			for _, out := range strings.Split(s[1], ",") {
				out = strings.ReplaceAll(out, " ", "")
				modules[name].AddOutput(out)
				inputs[out] = append(inputs[out], name)
			}
		case s[0] == "broadcaster":
			name := "broadcaster"

			modules[name] = NewBroadcast(name)

			for _, out := range strings.Split(s[1], ",") {
				out = strings.ReplaceAll(out, " ", "")
				modules[name].AddOutput(out)
				inputs[out] = append(inputs[out], name)
			}
		}
	}

	modules["button"] = NewButton("button")
	modules["rx"] = NewRX("rx")

	for name, inputs := range inputs {
		for _, i := range inputs {
			if m, ok := modules[name]; ok {
				m.AddInput(i)
			}
		}
	}

	fmt.Printf("modules =\n")
	for _, m := range modules {
		fmt.Printf("%v\n", m)
	}

	// Module(name=th, input=[xn qn xf zl], output=[rx])

	lookingForEvent = Event{from: "zl", to: "th", pulse: HIGH}
	eventOccurred = false

	firstSeen := 0
	secondSeen := 0

	buttonPressed := 0
	for {
		doRun()
		buttonPressed++
		if eventOccurred {
			if firstSeen == 0 {
				firstSeen = buttonPressed
				eventOccurred = false
			} else {
				secondSeen = buttonPressed
				break
			}
		}
	}

	fmt.Printf("firstSeen=%v\n", firstSeen)
	fmt.Printf("secondSeen=%v\n", secondSeen)
}

func areAllModulesInitialState() bool {
	for _, m := range modules {
		if !m.IsInitial() {
			return false
		}
	}
	return true
}

func doRun() {
	events := modules["button"].Process("", LOW)
	if debug {
		fmt.Printf("--------------------------------------------------------------------------------\n")
		fmt.Printf("initial events=%v\n", events)
		fmt.Printf("--------------------------------------------------------------------------------\n")
	}
	for len(events) > 0 {
		e := events[0]
		events = events[1:]

		if e.from == lookingForEvent.from && e.to == lookingForEvent.to && e.pulse == lookingForEvent.pulse {
			eventOccurred = true
		}

		if debug {
			fmt.Printf("%v -%v-> %v\n", e.from, e.pulse, e.to)
		}

		if m, ok := modules[e.to]; ok {
			res := m.Process(e.from, e.pulse)

			events = append(events, res...)
		}
	}
}
