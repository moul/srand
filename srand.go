package srand

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"errors"
	"os"
	"strconv"
	"time"
)

// Fast returns value that can be used to initialize random seed.
// That value is not cryptographically secure.
func Fast() int64 {
	return time.Now().UTC().UnixNano()
}

// Overridable will check if the "key" var is configured, else
// it will return a Fast random seed.
//
// If $SRAND is not parseable, panic is raised.
func Overridable(key string) (int64, error) {
	if env := os.Getenv(key); env != "" {
		n, err := strconv.ParseInt(env, 10, 64)
		if err != nil {
			return 0, err
		}
		return n, nil
	}
	return Fast(), nil
}

func MustOverridable(key string) int64 {
	ret, err := Overridable(key)
	if err != nil {
		panic(err)
	}
	return ret
}

// Secure returns a cryptographically secure random seed.
//
// Based on https://stackoverflow.com/a/54491783
func Secure() (int64, error) {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		return 0, errors.New("cannot seed math/rand package with cryptographically secure random number generator")
	}
	return int64(binary.LittleEndian.Uint64(b[:])), nil
}

func MustSecure() int64 {
	ret, err := Secure()
	if err != nil {
		panic(err)
	}
	return ret
}
