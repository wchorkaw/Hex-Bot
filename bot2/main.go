package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type command struct {
	command     string
	example     string
	description string
	numArgs     int
}

// Maps aren't ordered in go :sob:
var validCommandsOrder = []string{"init_board", "show_board", "make_move", "seto", "sety", "swap", "unset", "check_win", "quit"}
var validCommands = map[string]command{
	"init_board": {"init_board {digit}", "init_board 8", "Tells the bot to reset the game to an empty board w/ side length digit", 1},
	"show_board": {"show_board", "show_board", "Prints the board to stdout. Used for internal testing", 0},
	"make_move":  {"make_move", "make_move", "Asks the bot to give their move, based on the current board", 0},
	"seto":       {"seto {}", "seto a1", "Tells the bot about a move for the other bot", 1},
	"sety":       {"sety {}", "sety a1", "Tells the bot to play a move for itself", 1},
	"swap":       {"swap", "swap", "Uses the opening \"swap\" move in Hex", 0},
	"unset":      {"unset {}", "unset a1", "Tells the bot to set a tile as unused", 1},
	"check_win":  {"check_win", "check_win", "Tells the bot to check if the game is over. Returns 1 if itself has won, -1 if the opponent has won, 0 if the game has not terminated", 0},
	"quit":       {"quit", "quit", "The game is over", 0},
}

const unknownCommand = "Cmd not recognized. Please refer to known commands below:"
const wrongArgLen = "Not enough args. Please refer to the number of arguments below for the command:"

func printHelp(msg string) {
	fmt.Println(unknownCommand)
	fmt.Printf("%-30s%-30s%s\n", "Command", "Example", "Description")
	for _, c := range validCommandsOrder {
		v := validCommands[c]
		fmt.Printf("%-30s%-30s%s\n", v.command, v.example, v.description)
	}
	fmt.Println("\nNote that draws are impossible in hex, so no response for a draw is required")
}

func main() {
	bot := newBot(getColor())

	scanner := bufio.NewScanner(os.Stdin)
	// TODO: Run whatever computation is needed while waiting for input
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		args := []string{}
		for _, s := range strings.Split(line, " ") {
			args = append(args, strings.TrimSpace(s))
		}
		if args[0] == "quit" {
			break
		}
		if cmd, ok := validCommands[args[0]]; ok {
			if len(args)-1 < cmd.numArgs {
				printHelp(wrongArgLen)
				continue
			}
			err := bot.runCmd(args)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		} else {
			printHelp(unknownCommand)
		}
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
}
