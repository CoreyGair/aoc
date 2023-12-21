package y2023

import (
	"fmt"
	"slices"
	"strings"
)

func Day20(input string) (result int) {
	mods := ParseModules(input)

	lo, hi := 0, 0
	for i := 0; i < 1000; i++ {
		l, h := PressButton(mods)
		lo += l
		hi += h
	}
	return lo * hi
}

func Day20Part2(input string) (result int) {
	mods := ParseModules(input)

	var rxInput *Conjunction
	for name, mod := range mods {
		if slices.Contains(mod.GetDestinations(), "rx") {
			ok := false
			rxInput, ok = mods[name].(*Conjunction)
			if ok {
				break
			}
		}
	}

	rxInputCycleLength := make(map[string]int)
	for name := range rxInput.memory {
		rxInputCycleLength[name] = 0
	}

	i := 0
	for {
		i++
		{
			pulses := NewQueue[Pulse](100)

			pulses.Push(NewPulse("button", "broadcaster", false))

			for !pulses.Empty() {
				p := pulses.Pop()
				toMod, ok := mods[p.to]
				if !ok {
					continue
				}

				if p.to == rxInput.name {
					if p.hi && rxInputCycleLength[p.from] == 0 {
						rxInputCycleLength[p.from] = i
					}
				}

				pulses.PushAll(toMod.ProcessPulse(p))
			}
		}

		allDone := true
		for _, v := range rxInputCycleLength {
			if v == 0 {
				allDone = false
			}
		}
		if allDone {
			break
		}
	}

	xs := []int{}
	for _, v := range rxInputCycleLength {
		xs = append(xs, v)
	}
	return lcm(xs)
}

func ParseModules(input string) map[string]Module {
	lines := Lines(input)

	mods := make(map[string]Module)

	for _, l := range lines {
		newMod := ParseModule(l)
		mods[newMod.GetName()] = newMod
	}

	for name, mod := range mods {
		for _, destName := range mod.GetDestinations() {
			if destName == "output" {
				continue
			}
			if conj, ok := mods[destName].(*Conjunction); ok {
				conj.memory[name] = false
			}
		}
	}

	return mods
}

func PressButton(mods map[string]Module) (los int, his int) {
	pulses := NewQueue[Pulse](100)

	pulses.Push(NewPulse("button", "broadcaster", false))
	los++

	for !pulses.Empty() {
		p := pulses.Pop()
		toMod, ok := mods[p.to]
		if !ok {
			continue
		}

		newPulses := toMod.ProcessPulse(p)
		l, h := CountLoAndHi(newPulses)
		los += l
		his += h
		pulses.PushAll(newPulses)
	}
	return
}

func CountLoAndHi(pulses []Pulse) (los int, his int) {
	for _, p := range pulses {
		if p.hi {
			his++
		} else {
			los++
		}
	}
	return
}

type Pulse struct {
	from string
	to   string
	hi   bool
}

func NewPulse(from, to string, hi bool) Pulse {
	return Pulse{
		from: from,
		to:   to,
		hi:   hi,
	}
}

func (p Pulse) String() string {
	hi := "lo"
	if p.hi {
		hi = "hi"
	}
	return fmt.Sprintf("%s pulse from %s to %s", hi, p.from, p.to)
}

type ModuleCommon struct {
	name         string
	destinations []string
}

func (m ModuleCommon) GetName() string { return m.name }

func (m ModuleCommon) GetDestinations() []string { return m.destinations }

func (m ModuleCommon) SendPulses(hi bool) []Pulse {
	result := make([]Pulse, 0, len(m.destinations))
	for _, d := range m.destinations {
		result = append(result, NewPulse(m.name, d, hi))
	}
	return result
}

type Module interface {
	GetName() string
	GetDestinations() []string
	String() string
	ProcessPulse(pulse Pulse) []Pulse
}

func ParseModule(line string) Module {
	parts := strings.Split(strings.ReplaceAll(line, " ", ""), "->")

	moduleType := []rune(parts[0])[0]
	name := "broadcaster"
	if moduleType != 'b' {
		name = string([]rune(parts[0])[1:])
	}

	destinations := strings.Split(parts[1], ",")

	mc := ModuleCommon{
		name:         name,
		destinations: destinations,
	}

	switch moduleType {
	case 'b':
		return &Broadcaster{
			ModuleCommon: mc,
		}
	case '%':
		return &FlipFlop{
			ModuleCommon: mc,
			on:           false,
		}
	case '&':
		return &Conjunction{
			ModuleCommon: mc,
			memory:       make(map[string]bool),
		}
	default:
		panic("oop")
	}
}

type FlipFlop struct {
	ModuleCommon

	on bool
}

var _ Module = &FlipFlop{}

func (f *FlipFlop) ProcessPulse(pulse Pulse) []Pulse {
	if !pulse.hi {
		f.on = !f.on
		return f.SendPulses(f.on)
	}
	return []Pulse{}
}

func (f *FlipFlop) String() string {
	on := "off"
	if f.on {
		on = "on"
	}
	return fmt.Sprintf("flipflop %s %s destinations %s", f.name, on, strings.Join(f.destinations, ","))
}

type Conjunction struct {
	ModuleCommon

	memory map[string]bool
}

var _ Module = &Conjunction{}

func (c *Conjunction) ProcessPulse(pulse Pulse) []Pulse {
	c.memory[pulse.from] = pulse.hi

	for _, hi := range c.memory {
		if !hi {
			return c.SendPulses(true)
		}
	}
	return c.SendPulses(false)
}

func (c *Conjunction) String() string {
	sources := []string{}
	for k := range c.memory {
		sources = append(sources, k)
	}
	return fmt.Sprintf("conjunction %s sources: %s destinations: %s", c.name, strings.Join(sources, ","), strings.Join(c.destinations, ","))
}

type Broadcaster struct {
	ModuleCommon
}

var _ Module = &Broadcaster{}

func (b *Broadcaster) ProcessPulse(pulse Pulse) []Pulse {
	return b.SendPulses(pulse.hi)
}

func (b *Broadcaster) String() string {
	return fmt.Sprintf("broadcaster destinations: %s", strings.Join(b.destinations, ","))
}

type Queue[T any] struct {
	buf        []T
	start, end int
	empty      bool
}

func NewQueue[T any](cap int) Queue[T] {
	return Queue[T]{
		buf: make([]T, cap),
	}
}

func (q *Queue[T]) grow() {
	newBuf := make([]T, 2*len(q.buf))
	if q.end == q.start {
		q.start, q.end = 0, 0
		q.buf = newBuf
		return
	}

	if q.end < q.start {
		copy(newBuf, q.buf[q.start:])
		copy(newBuf, q.buf[0:q.end+1])
		q.end = len(q.buf) - q.start + q.end + 1
	} else {
		copy(newBuf, q.buf[q.start:q.end+1])
		q.end = q.end - q.start
	}
	q.start = 0
	q.buf = newBuf
}

func (q *Queue[T]) Push(t T) {
	if !q.empty && q.end == q.start {
		q.grow()
	}

	q.buf[q.end] = t
	q.end = (q.end + 1) % len(q.buf)
	q.empty = false
}

func (q *Queue[T]) Pop() T {
	t := q.buf[q.start]
	q.start = (q.start + 1) % len(q.buf)
	if q.start == q.end {
		q.empty = true
	}
	return t
}

func (q *Queue[T]) Empty() bool {
	return q.empty
}

func (q *Queue[T]) PushAll(ts []T) {
	for _, t := range ts {
		q.Push(t)
	}
}
