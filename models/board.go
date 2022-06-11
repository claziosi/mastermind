package models

import (
	"fmt"

	"naja.io/mastermind/constants"
)

type Board struct {
	BoardName string
	NbLines   int
	NbColumns int
	Board     [][]Peg
	Secret    []Peg
}

func (board Board) RevealSecret() {
	fmt.Println("Here is the secret code: ")
	for i := 0; i < constants.NbLines; i++ {
		fmt.Print(constants.Colour.String(board.Secret[i].Colour) + ", ")
	}
	fmt.Println()
}
