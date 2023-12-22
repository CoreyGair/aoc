package y2023

import (
	"os"
	"os/exec"
	"slices"
	"strconv"
	"time"
)

func Day17(input string) (result int) {
	costs := ParseCosts(input)

	path, states := FindPath(costs, 0, 0, len(costs)-1, len(costs[0])-1)

	for _, state := range states {
		s := make([]string, len(costs))
		for x := range s {
			for y := range costs[x] {
				if state.x == x && state.y == y {
					s[x] += "#"
				} else {
					s[x] += "."
				}
			}
		}
		PrintStringList(s)

		time.Sleep(time.Millisecond * 50)

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	return path[0].gcost
}

func ParseCosts(input string) [][]int {
	lines := Lines(input)

	costs := make([][]int, len(lines))
	for i, line := range lines {
		costs[i] = make([]int, 0, len(line))
		for _, c := range line {
			costs[i] = append(costs[i], Must1(strconv.Atoi(string(c))))
		}
	}
	return costs
}

type State struct {
	fcost int
	gcost int

	prev *State

	x, y   int
	dx, dy int
}

func NewState(fcost, gcost, x, y, dx, dy int) *State {
	return &State{
		fcost: fcost,
		gcost: gcost,
		prev:  nil,
		x:     x,
		y:     y,
		dx:    dx,
		dy:    dy,
	}
}

func (s *State) Priority() int {
	return s.fcost
}

func (s *State) Neighbours(costs [][]int, h func(x, y int) int) (ns []*State) {
	if s.dx == 0 && s.dy == 0 {
		{
			x := s.x + 1
			y := s.y
			g := s.gcost + costs[x][y]
			ns = append(ns, NewState(g+h(x, y), g, x, y, 1, 0))
		}
		{
			x := s.x
			y := s.y + 1
			g := s.gcost + costs[x][y]
			ns = append(ns, NewState(g+h(x, y), g, x, y, 0, 1))
		}
		return
	}

	if s.dx != 0 {
		if AbsInt(s.dx) < 3 {
			dx := 1
			if s.dx < 1 {
				dx = -1
			}

			x := s.x + dx
			if x >= 0 && x < len(costs) {
				g := s.gcost + costs[x][s.y]
				ns = append(ns, NewState(g+h(x, s.y), g, x, s.y, s.dx+dx, 0))
			}
		}
		if s.y+1 < len(costs[0]) {
			g := s.gcost + costs[s.x][s.y+1]
			ns = append(ns, NewState(g+h(s.x, s.y+1), g, s.x, s.y+1, 0, 1))
		}
		if s.y-1 >= 0 {
			g := s.gcost + costs[s.x][s.y-1]
			ns = append(ns, NewState(g+h(s.x, s.y-1), g, s.x, s.y-1, 0, -1))
		}
	}

	if s.dy != 0 {
		if AbsInt(s.dy) < 3 {
			dy := 1
			if s.dy < 1 {
				dy = -1
			}

			y := s.y + dy
			if y >= 0 && y < len(costs[0]) {
				g := s.gcost + costs[s.x][y]
				ns = append(ns, NewState(g+h(s.x, y), g, s.x, y, 0, s.dy+dy))
			}
		}
		if s.x+1 < len(costs) {
			g := s.gcost + costs[s.x+1][s.y]
			ns = append(ns, NewState(g+h(s.x+1, s.y), g, s.x+1, s.y, 1, 0))
		}
		if s.x-1 >= 0 {
			g := s.gcost + costs[s.x-1][s.y]
			ns = append(ns, NewState(g+h(s.x-1, s.y), g, s.x-1, s.y, -1, 0))
		}
	}

	return
}

func (s *State) Hash() int {
	return s.x*1000000 + s.y*100 + s.dx*10 + s.dy
}

func FindPath(costs [][]int, startX, startY, goalX, goalY int) ([]*State, []*State) {
	h := func(x, y int) int { return AbsInt(x-goalX) + AbsInt(y-goalY) }

	openSet := NewPriorityQueue[*State](func(t1, t2 *State) bool { return t1.x == t2.x && t1.y == t2.y && t1.dx == t2.dx && t1.dy == t2.dy })
	openSet.Push(NewState(h(startX, startY), 0, 0, 0, 0, 0))

	seen := make(map[int]struct{})

	states := make([]*State, 0)

	for openSet.Length() > 0 {
		currentState := openSet.Pop()
		seen[currentState.Hash()] = struct{}{}

		states = append(states, currentState)
		if len(states) == 10000 {
			return nil, states
		}

		if currentState.x == goalX && currentState.y == goalY {
			path := make([]*State, 0)
			for currentState != nil {
				path = append(path, currentState)
				currentState = currentState.prev
			}
			return path, states
		}

		for _, newState := range currentState.Neighbours(costs, h) {
			newState.prev = currentState

			_, haveSeen := seen[newState.Hash()]

			if oldNewState, contains := openSet.Contains(newState); !haveSeen && contains && newState.gcost < oldNewState.gcost {
				openSet.Remove(oldNewState)
				openSet.Push(newState)
				continue
			}
			if !haveSeen {
				openSet.Push(newState)
				continue
			}
		}
	}
	return []*State{}, states
}

type Priority interface {
	Priority() int
}

type PriorityQueue[T Priority] struct {
	q  []T
	eq func(t1, t2 T) bool
}

func NewPriorityQueue[T Priority](eq func(t1, t2 T) bool) PriorityQueue[T] {
	return PriorityQueue[T]{
		q:  make([]T, 0),
		eq: eq,
	}
}

func (pq *PriorityQueue[T]) Length() int {
	return len(pq.q)
}

func (pq *PriorityQueue[T]) Push(t T) {
	tPriority := t.Priority()
	for i, t2 := range pq.q {
		if tPriority < t2.Priority() {
			pq.q = slices.Insert(pq.q, i, t)
			return
		}
	}
	pq.q = append(pq.q, t)
}

func (pq *PriorityQueue[T]) Pop() T {
	t := pq.q[0]
	pq.q = pq.q[1:]
	return t
}

func (pq *PriorityQueue[T]) Contains(t T) (T, bool) {
	var empty T
	for _, t2 := range pq.q {
		if pq.eq(t, t2) {
			return t2, true
		}
		if t2.Priority() > t.Priority() {
			return empty, false
		}
	}
	return empty, false
}

func (pq *PriorityQueue[T]) Remove(t T) {
	for i, t2 := range pq.q {
		if pq.eq(t, t2) {
			pq.q = slices.Delete(pq.q, i, i+1)
		}
	}
}
