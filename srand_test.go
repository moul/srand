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

func ExampleMustOverridable() {
	rand.Seed(srand.MustOverridable("SRAND")) // seed with Fast

	os.Setenv("SRAND", "42")
	rand.Seed(srand.MustOverridable("SRAND")) // seed with 42
}

func TestMustOverridable(t *testing.T) {
	a := srand.Fast()
	b := srand.Fast()
	if a >= b {
		t.Errorf("Expected %d > %d.", b, a)
	}

	os.Setenv("SRAND", "42")
	c, err := srand.Overridable("SRAND")
	if err != nil {
		t.Errorf("err should be nil: %v", err)
	}
	if c != 42 {
		t.Errorf("Expected 42, got %d.", c)
	}
}

func ExampleMustSecure() {
	rand.Seed(srand.MustSecure())
}

func TestSecure(t *testing.T) {
	a, err := srand.Secure()
	if err != nil {
		t.Errorf("err should be nil: %v", err)
	}
	b, err := srand.Secure()
	if err != nil {
		t.Errorf("err should be nil: %v", err)
	}
	if a == b {
		t.Errorf("Expected %d != %d.", b, a)
	}
}
