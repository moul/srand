package srand

import (
	"os"
	"strconv"
	"time"
)

// Fast returns value that can be used to initialize random seed.
// That value is not cryptographically secure.
func Fast() int64 {
	return time.Now().UTC().UnixNano()
}

// Overridable will check if the $SRAND var is configured, else
// it will return a Fast random seed.
//
// If $SRAND is not parseable, panic is raised.
func Overridable() int64 {
	if env := os.Getenv("SRAND"); env != "" {
		n, err := strconv.ParseInt(env, 10, 64)
		if err != nil {
			panic(err)
		}
		return n
	}
	return Fast()
}