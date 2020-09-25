package srand_test

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
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

func ExampleSafeFast() {
	go func() { rng := rand.New(rand.NewSource(srand.SafeFast())); fmt.Println(rng.Intn(42)) }()
	go func() { rng := rand.New(rand.NewSource(srand.SafeFast())); fmt.Println(rng.Intn(42)) }()
	go func() { rng := rand.New(rand.NewSource(srand.SafeFast())); fmt.Println(rng.Intn(42)) }()
	go func() { rng := rand.New(rand.NewSource(srand.SafeFast())); fmt.Println(rng.Intn(42)) }()
	go func() { rng := rand.New(rand.NewSource(srand.SafeFast())); fmt.Println(rng.Intn(42)) }()
}

func TestConcurrentFast(t *testing.T) {
	ch := make(chan int64, 100)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			ch <- srand.Fast()
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	checkMap := map[int64]bool{}
	for result := range ch {
		checkMap[result] = true
	}

	if len(checkMap) != 100 {
		t.Errorf("srand.Fast is not thread safe.")
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
