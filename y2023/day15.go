package y2023

import (
	"regexp"
	"strconv"
	"strings"
)

func Day15(input string) (result int) {
	parts := strings.Split(strings.ReplaceAll(input, "\n", ""), ",")

	for _, part := range parts {
		result += HASH(part)
	}
	return
}

func Day15Part2(input string) (result int) {
	parts := strings.Split(strings.ReplaceAll(input, "\n", ""), ",")

	boxes := make([]Box, 256)
	for i := range boxes {
		boxes[i] = NewBox()
	}

	steps := make([]Step, 0, len(parts))
	for _, part := range parts {
		steps = append(steps, ParseStep(part))
	}

	for _, step := range steps {
		boxIndex := step.labelHash

		if step.operation == Remove {
			boxes[boxIndex].RemoveLensWithLabel(step.label)
		} else if step.operation == Set {
			boxes[boxIndex].Set(step.Lens())
		} else {
			panic("oops")
		}
	}

	for i, box := range boxes {
		result += box.SumFocusingPowers(i)
	}

	return
}

func HASH(s string) (result int) {
	for _, c := range s {
		result = ((result + int(c)) * 17) % 256
	}
	return
}

type Op rune

const (
	Remove Op = '-'
	Set    Op = '='
)

type Step struct {
	label       string
	labelHash   int
	operation   Op
	focalLength int
}

func ParseStep(input string) Step {
	label := regexp.MustCompile("[a-z]+").FindString(input)
	op := Op([]rune(regexp.MustCompile("-|=").FindString(input))[0])
	focalLength := Must1(strconv.Atoi(regexp.MustCompile("[1-9]").FindString(input)))

	return Step{
		label:       label,
		labelHash:   HASH(label),
		operation:   op,
		focalLength: focalLength,
	}
}

func (s *Step) Lens() Lens {
	return Lens{
		focalLength: s.focalLength,
		label:       s.label,
	}
}

type Lens struct {
	focalLength int
	label       string
}

type Box struct {
	lenses []Lens
}

func NewBox() Box {
	return Box{
		lenses: make([]Lens, 0),
	}
}

func (b *Box) RemoveLensWithLabel(label string) {
	newLenses := make([]Lens, 0, len(b.lenses))
	for _, lens := range b.lenses {
		if lens.label == label {
			continue
		}
		newLenses = append(newLenses, lens)
	}
	b.lenses = newLenses
	return
}

func (b *Box) Set(newLens Lens) {
	done := false
	newLenses := make([]Lens, 0, len(b.lenses)+1)
	for _, lens := range b.lenses {
		if lens.label == newLens.label {
			lens.focalLength = newLens.focalLength
			done = true
		}
		newLenses = append(newLenses, lens)
	}
	if !done {
		newLenses = append(newLenses, newLens)
	}
	b.lenses = newLenses
	return
}

func (b *Box) SumFocusingPowers(boxIndex int) (result int) {
	for i, lens := range b.lenses {
		result += (1 + boxIndex) * (1 + i) * lens.focalLength
	}
	return
}
