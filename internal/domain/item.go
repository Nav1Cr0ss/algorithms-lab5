package domain

import (
	"math/rand"
	"time"
)

type Item struct {
	Value  int
	Weight int
}

type Individual struct {
	Genes   []bool
	Fitness int
}

var (
	randSource    = rand.NewSource(time.Now().UnixNano())
	randGenerator = rand.New(randSource)
)

func NewRandItem() Item {
	return Item{
		Value:  randGenerator.Intn(19) + 2,
		Weight: randGenerator.Intn(10) + 1,
	}
}
