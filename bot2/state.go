package main

import (
	"fmt"
	"strconv"
	"strings"
)

type state struct {
	board     board
	boardSize int
	groups    map[color]*unionFind
}

type board [][]color

func (b board) get(c cell) color       { return b[c.row][c.col] }
func (b *board) set(c cell, val color) { (*b)[c.row][c.col] = val }

// Custom error for placing on board {{{

type cellOccupied struct{}

func (c cellOccupied) Error() string { return "cell occupied" }

// }}}

// Some constants that aren't constants {{{

var edge1 = cell{999, 999}
var edge2 = cell{998, 998}
var opp = map[color]color{white: black, black: white, empty: empty}
var cellOccupiedError = cellOccupied{}
var neighborPatterns = []cell{{-1, 0}, {0, -1}, {-1, 1}, {0, 1}, {1, 0}, {1, -1}}

// }}}

func newBoard(size int) *state {
	s := state{
		board:     make([][]color, 0, size),
		boardSize: size,
		groups:    map[color]*unionFind{white: newUnionFind(), black: newUnionFind()},
	}

	for i := 0; i < size; i++ {
		arr := make([]color, size)
		for i := 0; i < size; i++ {
			arr[i] = empty
		}
		s.board = append(s.board, arr)
	}
	return &s
}

func (s *state) checkWin() color {
	if s.groups[white].connected(edge1, edge2) {
		return white
	} else if s.groups[black].connected(edge1, edge2) {
		return black
	} else {
		return empty
	}
}

func (s *state) placePiece(cell cell, color color) error {
	if s.board.get(cell) == empty {
		s.board.set(cell, color)
	} else {
		return cellOccupiedError
	}

	// TODO: Break up these union finds
	// if we empty a cell that used to be in a group
	if color != empty {
		if cell.edge(color) == 0 {
			s.groups[color].join(edge1, cell)
		} else if cell.edge(color) == s.boardSize-1 {
			s.groups[color].join(edge2, cell)
		}
		s.joinPieces(cell, color)
	}

	return nil
}

func (s *state) joinPieces(c cell, color color) {
	x, y, size := c.row, c.col, s.boardSize
	for _, neighbor := range neighborPatterns {
		n := cell{neighbor.row + x, neighbor.col + y}
		if 0 <= n.row && n.row < size &&
			0 <= n.col && n.col < size &&
			s.board.get(n) == color {
			s.groups[color].join(n, c)
		}
	}
}

type paddedBuilder struct{ strings.Builder }

func (b *paddedBuilder) WriteRune(r rune, num int) {
	b.Builder.WriteRune(r)
	b.pad(num)
}

func (b *paddedBuilder) WriteString(s string, num int) {
	b.Builder.WriteString(s)
	b.pad(num)
}

func (b *paddedBuilder) pad(num int) { b.Builder.WriteString(strings.Repeat(" ", num)) }

func (s *state) String() string {
	tileChars := map[color]rune{
		empty: '.',
		black: 'B',
		white: 'W',
	}

	var sb paddedBuilder
	size := len(strconv.Itoa(s.boardSize))
	offset := 1
	sb.pad(offset + 1)
	for x := 0; x < s.boardSize; x++ {
		sb.WriteRune(rune(int('A')+x), offset*2)
	}
	fmt.Fprintln(&sb)
	for y := 0; y < s.boardSize; y++ {
		num := strconv.Itoa(y + 1)
		sb.WriteString(num, offset*2+size-len(num))
		for x := 0; x < s.boardSize; x++ {
			sb.WriteRune(tileChars[s.board[y][x]], offset*2)
		}
		sb.Builder.WriteRune(tileChars[white])
		sb.WriteString("\n", offset*(y+1))
	}
	sb.pad(offset*2 + 1)
	sb.Builder.WriteString(strings.Repeat(string(tileChars[black])+strings.Repeat(" ", offset*2), s.boardSize))

	return sb.String()
}
