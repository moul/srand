package srand

import "time"

// Fast returns value that can be used to initialize random seed.
// That value is not cryptographically secure.
func Fast() int64 {
	return time.Now().UTC().UnixNano()
}
