package y2023

import (
	"regexp"
	"slices"
	"strings"
)

type Direction = rune

const (
	Right Direction = 'R'
	Left  Direction = 'L'
)

type Node string

var (
	nodeRegex = regexp.MustCompile("[A-Z0-9]{3}")
)

type DesertMapElement struct {
	source Node
	left   *DesertMapElement
	right  *DesertMapElement
}

func ParseDesertMapElements(lines []string) []DesertMapElement {
	elems := make([]DesertMapElement, 0)
	destLefts := make([]Node, 0)
	destRights := make([]Node, 0)

	for _, l := range lines {
		if l == "" {
			continue
		}

		nodes := nodeRegex.FindAllString(l, -1)
		elems = append(elems, DesertMapElement{
			source: Node(nodes[0]),
			left:   nil,
			right:  nil,
		})
		destLefts = append(destLefts, Node(nodes[1]))
		destRights = append(destRights, Node(nodes[2]))
	}

	for i := range elems {
		dl := destLefts[i]
		dr := destRights[i]
		for j := range elems {
			m := &elems[j]
			if m.source == dl {
				elems[i].left = m
			}
			if m.source == dr {
				elems[i].right = m
			}
		}
	}

	return elems
}

func (m *DesertMapElement) Step(instr Direction) *DesertMapElement {
	if instr == Left {
		return m.left
	}
	return m.right
}

func Day8(input string) int {
	lines := strings.Split(input, "\n")

	instructions := lines[0]

	mapElements := ParseDesertMapElements(lines[1:])

	steps := 0
	currNode := &mapElements[slices.IndexFunc(mapElements, func(dme DesertMapElement) bool { return dme.source == "AAA" })]
	for {
		for _, instr := range instructions {
			currNode = currNode.Step(instr)
			steps++
			if currNode.source == "ZZZ" {
				return steps
			}
		}
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(xs []int) (ans int) {
	ans = xs[0]
	for i := 1; i < len(xs); i++ {
		ans = ((xs[i] * ans) / (gcd(xs[i], ans)))
	}
	return
}

func Day8Part2(input string) int {
	lines := strings.Split(input, "\n")

	instructions := lines[0]

	mapElements := ParseDesertMapElements(lines[1:])

	var startNodes []*DesertMapElement
	for i, m := range mapElements {
		if strings.HasSuffix(string(m.source), "A") {
			startNodes = append(startNodes, &mapElements[i])
		}
	}

	distances := make([]int, len(startNodes))
	for i := range startNodes {
		currNode := startNodes[i]
		distances[i] = 0
		for !strings.HasSuffix(string(currNode.source), "Z") {
			for _, instr := range instructions {
				currNode = currNode.Step(instr)
				distances[i]++
				if strings.HasSuffix(string(currNode.source), "Z") {
					break
				}
			}
		}
	}

	return lcm(distances)
}
