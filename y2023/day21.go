package y2023

func Day21Tester(input string) int {
	rows := Lines(input)

	start := FindStart(rows)

	reachable := make(HashSet[Location])
	reachable.Insert(start)

	reachable = TakeSteps(reachable, rows, 6)

	return len(reachable)
}

func Day21(input string) int {
	rows := Lines(input)

	start := FindStart(rows)

	reachable := make(HashSet[Location])
	reachable.Insert(start)

	reachable = TakeSteps(reachable, rows, 64)

	return len(reachable)
}

func Day21Part2(input string) int {
	copies := 5

	rows := Lines(input)

	steps := []int{}
	for i := 0; i < 3; i++ {
		steps = append(steps, (len(rows)/2)+(len(rows)*i))
	}

	newRows := make([]string, len(rows)*copies)
	for i := 0; i < len(newRows); i++ {
		newRows[i] = ""
		for j := 0; j < copies; j++ {
			newRows[i] += rows[i%len(rows)]
		}
	}
	for i, row := range newRows {
		for j, letter := range row {
			if letter == 'S' && (i != len(newRows)/2 || j != len(newRows[0])/2) {
				newRows[i] = replaceAtNoCopy(newRows[i], j, '.')
			}
		}
	}

	rows = newRows

	start := FindStart(rows)

	reachable := make(HashSet[Location])
	reachable.Insert(start)

	// run for each of steps[0,1,2],
	// fit polynomial and extrapolate to x=202300
	reachable = TakeSteps(reachable, rows, steps[0])

	return len(reachable)
}

func TakeSteps(reachable HashSet[Location], rows []string, steps int) HashSet[Location] {
	for i := 0; i < steps; i++ {
		newReachable := make(HashSet[Location])

		for _, prevReachable := range reachable {
			for _, n := range prevReachable.Neighbours() {
				if n.row < 0 {
					n.row += len(rows)
				}
				if n.col < 0 {
					n.col += len(rows[0])
				}
				n.row %= len(rows)
				n.col %= len(rows[0])
				if []rune(rows[n.row])[n.col] == '#' {
					continue
				}
				newReachable.Insert(n)
			}
		}

		if newReachable.Equal(reachable) {
			return newReachable
		}

		reachable = newReachable
	}
	return reachable
}

func FindStart(rows []string) Location {
	for r, row := range rows {
		for c, letter := range row {
			if letter == 'S' {
				return Location{
					row: r,
					col: c,
				}
			}
		}
	}
	panic("")
}

func (l Location) Hash() int {
	return l.row*1000000 + l.col
}

type Hashable interface {
	Hash() int
}

type HashSet[T Hashable] map[int]T

func (hs *HashSet[T]) Insert(t T) {
	(*hs)[t.Hash()] = t
}

func (hs *HashSet[T]) Remove(t T) {
	delete(*hs, t.Hash())
}

func (hs *HashSet[T]) Contains(t T) bool {
	_, ok := (*hs)[t.Hash()]
	return ok
}

func (hs *HashSet[T]) Equal(other HashSet[T]) bool {
	for k, _ := range *hs {
		if _, ok := (other)[k]; !ok {
			return false
		}
	}
	return true
}

func (l Location) Neighbours() []Location {
	return []Location{
		{row: l.row + 1, col: l.col},
		{row: l.row - 1, col: l.col},
		{row: l.row, col: l.col + 1},
		{row: l.row, col: l.col - 1},
	}
}
