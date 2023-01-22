package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akamensky/argparse"
)

type color int

const (
	empty color = 10
	white color = 3
	black color = 4
)

func getColor() color {
	parser := argparse.NewParser("hexbot", "Example hex bot that makes random valid placements")
	color := parser.StringPositional(&argparse.Options{Required: true, Help: "This bot's color. White is left->right"})
	err := parser.Parse(os.Args)
	if err != nil {
		log.Fatal(parser.Usage(err))
	}
	if *color == "white" {
		return white
	} else if *color == "black" {
		return black
	} else {
		log.Fatal(parser.Usage(fmt.Errorf("color must be 'white' or 'black'")))
		return empty
	}
}
