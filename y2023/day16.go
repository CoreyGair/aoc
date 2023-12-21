package y2023

import (
	"maps"
	"slices"
	"strings"
)

func Day16(input string) (result int) {
	lines := Lines(input)

	tiles := make([][]Tile, len(lines))
	for i, row := range lines {
		tiles[i] = make([]Tile, 0, len(row))
		for _, letter := range row {
			tiles[i] = append(tiles[i], NewTile(TileType(letter)))
		}
	}

	Energize(tiles, Location{row: 0, col: 0}, East)

	for _, r := range tiles {
		for _, t := range r {
			if t.Energized {
				result++
			}
		}
	}
	return
}

func Day16Part2(input string) (result int) {
	lines := Lines(input)

	tiles := make([][]Tile, len(lines))
	for i, row := range lines {
		tiles[i] = make([]Tile, 0, len(row))
		for _, letter := range row {
			tiles[i] = append(tiles[i], NewTile(TileType(letter)))
		}
	}

	for i := range tiles {
		pos := Location{row: i, col: 0}
		dir := East
		x := Energize2(tiles, pos, dir)
		if x > result {
			result = x
		}

		pos = Location{row: i, col: len(tiles[0]) - 1}
		dir = West
		x = Energize2(tiles, pos, dir)
		if x > result {
			result = x
		}
	}
	for i := range tiles[0] {
		pos := Location{row: 0, col: i}
		dir := South
		x := Energize2(tiles, pos, dir)
		if x > result {
			result = x
		}

		pos = Location{row: len(tiles) - 1, col: i}
		dir = North
		x = Energize2(tiles, pos, dir)
		if x > result {
			result = x
		}
	}
	return
}

func Energize2(tiles [][]Tile, pos Location, dir Cardinal) (result int) {
	Energize(tiles, pos, dir)

	for i, r := range tiles {
		for j, t := range r {
			if t.Energized {
				result++
			}
			tiles[i][j].Reset()
		}
	}
	return
}

func Lines(input string) []string {
	return FilterNewlines(strings.Split(input, "\n"))
}

type TileType rune

const (
	Empty           TileType = '.'
	MirrorUp        TileType = '/'
	MirrorDown      TileType = '\\'
	SplitVertical   TileType = '|'
	SplitHorizontal TileType = '-'
)

var (
	identity = map[Cardinal][]Cardinal{
		North: {North},
		East:  {East},
		South: {South},
		West:  {West},
	}

	tileTypeToNextDir = map[TileType]map[Cardinal][]Cardinal{
		Empty: identity,
		SplitVertical: With(identity, map[Cardinal][]Cardinal{
			East: {North, South},
			West: {North, South},
		}),
		SplitHorizontal: With(identity, map[Cardinal][]Cardinal{
			North: {West, East},
			South: {West, East},
		}),
		MirrorDown: {
			North: {West},
			East:  {South},
			South: {East},
			West:  {North},
		},
		MirrorUp: {
			North: {East},
			East:  {North},
			South: {West},
			West:  {South},
		},
	}
)

func With(m map[Cardinal][]Cardinal, n map[Cardinal][]Cardinal) map[Cardinal][]Cardinal {
	new := maps.Clone(m)
	for k, v := range n {
		new[k] = v
	}
	return new
}

type Tile struct {
	Type           TileType
	Energized      bool
	SeenDirections []Cardinal
}

func NewTile(t TileType) Tile {
	return Tile{
		Type:           t,
		Energized:      false,
		SeenDirections: []Cardinal{},
	}
}

func (t *Tile) Reset() {
	t.Energized = false
	t.SeenDirections = []Cardinal{}
}

func (t *Tile) Energize(inDirection Cardinal) []Cardinal {
	if slices.Contains(t.SeenDirections, inDirection) {
		return []Cardinal{}
	}
	t.Energized = true
	t.SeenDirections = append(t.SeenDirections, inDirection)
	return tileTypeToNextDir[t.Type][inDirection]
}

func Energize(tiles [][]Tile, pos Location, direction Cardinal) {
	if pos.row < 0 || pos.row >= len(tiles) || pos.col < 0 || pos.col >= len(tiles[0]) {
		return
	}

	nextDirs := tiles[pos.row][pos.col].Energize(direction)
	for _, d := range nextDirs {
		Energize(tiles, d.Move(pos), d)
	}
}
