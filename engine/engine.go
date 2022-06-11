package engine

import (
	"math/rand"
	"time"

	"naja.io/mastermind/constants"
	"naja.io/mastermind/models"
)

// Function accepting arguments
func RandomSecret() []models.Peg {

	rand.Seed(time.Now().UnixNano())

	var pegs []models.Peg

	for i := 0; i < constants.NbLines; i++ {
		peg := models.Peg{
			Colour: constants.Colour(rand.Intn(9)),
		}
		pegs = append(pegs, peg)
	}
	return pegs
}
