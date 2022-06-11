package main

import (
	"fmt"

	"naja.io/mastermind/constants"
	"naja.io/mastermind/engine"
	"naja.io/mastermind/models"
)

func main() {
	fmt.Println("Welcome to Mastermind in Golang !")

	myBoard := models.Board{
		BoardName: "MyBoard",
		NbLines:   constants.NbLines,
		NbColumns: constants.NbColumns,
		Secret:    engine.RandomSecret(),
	}

	myBoard.RevealSecret()

}
