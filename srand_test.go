package srand_test

import (
	"math/rand"
	"os"
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
		t.Errorf("Expected %d > %d.", b, a)
	}
}

func ExampleOverridable() {
	rand.Seed(srand.Overridable()) // seed with Fast

	os.Setenv("SRAND", "42")
	rand.Seed(srand.Overridable()) // seed with 42
}

func TestOverridable(t *testing.T) {
	a := srand.Fast()
	b := srand.Fast()
	if a >= b {
		t.Errorf("Expected %d > %d.", b, a)
	}

	os.Setenv("SRAND", "42")
	c := srand.Overridable()
	if c != 42 {
		t.Errorf("Expected 42, got %d.", c)
	}
}
