package main

import (
	"fmt"

	"./models"
)

func main() {
	fmt.Println("hello world")

	e := models.Board{
		BoardName: "MyBoard",
		NbLines:   10,
		NbColumns: 4,
	}
	e.LeavesRemaining()

}
