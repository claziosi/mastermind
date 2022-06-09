package models

import "fmt"

type Board struct {
	BoardName string
	NbLines   int
	NbColumns int
}

func (e Board) LeavesRemaining() {
	fmt.Printf("%s has %d lines and %d columns\n", e.BoardName, e.NbLines, e.NbColumns)
}
