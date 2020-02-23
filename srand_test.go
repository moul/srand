package srand_test

import (
	"math/rand"
	"testing"

	"moul.io/srand"
)

func ExampleFast() {
	rand.Seed(srand.Fast())
}

func TestFast(t *testing.T) {
	a := srand.Fast()
	b := srand.Fast()
	if a >= b {
		t.Fatalf("Expected %d > %d.", b, a)
	}
}
