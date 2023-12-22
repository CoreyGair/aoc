package y2023

import (
	"fmt"
	"slices"
)

func Day22(input string) int {
	bricks := ParseBrickList(input)
	slices.SortFunc[[]Brick, Brick](bricks, ZSortBricks)

	supportedBy := make(map[int][]int)
	for i := range bricks {
		supportedBy[i] = make([]int, 0, 5)
	}
	supports := make(map[int][]int)
	for i := range bricks {
		supports[i] = make([]int, 0, 5)
	}

	for i, b := range bricks {
		for j := i + 1; j < len(bricks); j++ {
			b2 := bricks[j]
			if b2.ends[0].z <= b.ends[1].z {
				continue
			}

			if b.Overlaps(b2) {
				fmt.Printf("%s maybe supports %s\n", b, b2)
				supported := true
				for _, b3i := range supports[i] {
					b3 := bricks[b3i]
					if b2.Overlaps(b3) {
						fmt.Printf("but %s suported by %s and overlaps %s so no support\n", b3, b, b2)
						supported = false
						break
					}
				}
				if supported {
					supports[i] = append(supports[i], j)
					supportedBy[j] = append(supportedBy[j], i)
				}
			}
		}
	}

	result := 0
	for i := range bricks {
		if len(supports[i]) == 0 {
			result++
			continue
		}

		allSuportedHaveOtherSupports := true
		for j := range supports[i] {
			if len(supportedBy[j]) == 1 {
				allSuportedHaveOtherSupports = false
				break
			}
		}
		if allSuportedHaveOtherSupports {
			result++
			continue
		}
	}

	return result
}

type Point struct {
	x, y, z int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d,%d,%d)", p.x, p.y, p.z)
}

func ZSort(a, b Point) int {
	return a.z - b.z
}

type Brick struct {
	ends [2]Point
}

func (b Brick) String() string {
	return fmt.Sprintf("{%s, %s}", b.ends[0], b.ends[1])
}

func ZSortBricks(a, b Brick) int {
	return ZSort(a.ends[0], b.ends[0])
}

func ParseBrick(row string) Brick {
	list := parseIntSlice(row)

	b := Brick{[2]Point{{list[0], list[1], list[2]}, {list[3], list[4], list[5]}}}

	if b.ends[1].z < b.ends[0].z {
		tmp := b.ends[0]
		b.ends[0] = b.ends[1]
		b.ends[1] = tmp
	}
	return b
}

func ParseBrickList(input string) (bricks []Brick) {
	lines := Lines(input)
	bricks = make([]Brick, 0, len(lines))
	for _, line := range lines {
		bricks = append(bricks, ParseBrick(line))
	}
	return
}

func (b Brick) Overlaps(other Brick) bool {
	return OneDimOverlaps(b.ends[0].x, b.ends[1].x, other.ends[0].x, other.ends[1].x) && OneDimOverlaps(b.ends[0].y, b.ends[1].y, other.ends[0].y, other.ends[1].y)
}

func OneDimOverlaps(a1, a2, b1, b2 int) bool {
	aMax, aMin := a1, a2
	if aMax < aMin {
		aMax, aMin = aMin, aMax
	}
	bMax, bMin := b1, b2
	if bMax < bMin {
		bMax, bMin = bMin, bMax
	}

	return (aMin <= bMin && aMax >= bMin) || (aMax >= bMax && aMin <= bMax) || (aMax <= bMax && aMin >= bMin) || (bMin >= aMin && bMax <= aMax)
}

func ccw(a, b, c Point) bool {
	return (c.y-a.y)*(b.x-a.x) > (b.y-a.y)*(c.x-a.x)
}

func intersects(A, B [2]Point) bool {
	return ccw(A[0], B[0], B[1]) != ccw(A[1], B[0], B[1]) && ccw(A[0], A[1], B[0]) != ccw(A[0], A[1], B[1])
}
