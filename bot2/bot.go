package main

import (
	"fmt"
	"log"
	"strconv"
)

type randomHexBot struct {
	color     color
	opp       color
	moveCount int
	state     *state
}

func newBot(color color) randomHexBot {
	bot := randomHexBot{
		color:     color,
		opp:       opp[color],
		moveCount: 0,
		state:     newBoard(11),
	}

	return bot
}

func (b *randomHexBot) initBoard(size string) error {
	s, err := strconv.Atoi(size)
	if err != nil {
		return err
	}
	b.state = newBoard(s)
	return nil
}

func (b *randomHexBot) showBoard() error {
	fmt.Println(b.state)
	return nil
}

func (b *randomHexBot) makeMove() error { return nil }

type winner int

const (
	botWin winner = 1
	oppWin winner = -1
	noWin  winner = 0
)

// TODO: Handle swap move
func (b *randomHexBot) swapMove() error { return nil }
func (b *randomHexBot) getWinner() winner {
	color := b.state.checkWin()
	if color == b.color {
		return botWin
	} else if color == b.opp {
		return oppWin
	} else {
		return noWin
	}
}

func (b *randomHexBot) setCell(move string, color color) error {
	cell, err := moveToCell(move, b.state.boardSize)
	if err != nil {
		return err
	}
	return b.state.placePiece(cell, color)
}
func (b *randomHexBot) playOtherBot(move string) error { return b.setCell(move, b.opp) }
func (b *randomHexBot) playSelf(move string) error     { return b.setCell(move, b.color) }
func (b *randomHexBot) unsetTile(move string) error    { return b.setCell(move, empty) }

func (b *randomHexBot) runCmd(args []string) error {
	switch args[0] {
	case "init_board":
		return b.initBoard(args[1])
	case "show_board":
		return b.showBoard()
	case "make_move":
		return b.makeMove()

	case "seto":
		return b.playOtherBot(args[1])
	case "sety":
		return b.playSelf(args[1])
	case "unset":
		return b.unsetTile(args[1])
	case "swap":
		return b.swapMove()
	case "check_win":
		fmt.Printf("%d\n", b.getWinner())
	default:
		log.Fatalf("unhandled valid command: %s", args[0])
	}
	return nil
}

// TODO: Add error handling?
func moveToCell(move string, size int) (cell, error) {
	row, err := strconv.Atoi(move[1:])
	if err != nil {
		return cell{-1, -1}, nil
	}
	col := int(move[0]) - int('a')
	return cell{row: row - 1, col: col}, nil
}
