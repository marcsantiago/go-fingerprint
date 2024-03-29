package fingerprint

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
	"unsafe"

	"github.com/minio/highwayhash"
)

var (
	_key  = make([]byte, 32)
	_seed = rand.NewSource(time.Now().Unix())
)

func init() {
	r := rand.New(_seed)
	_, _ = r.Read(_key)
}

var dashes = [4]int{8, 13, 18, 23}

// Identify takes the string representation of objects and creates a hash in the form of a uuid
// if a seed is not provided a random 32 byte key is generated using a timestamp
func Identify(strObjects ...fmt.Stringer) (string, error) {
	hh, err := highwayhash.New128(_key)
	if err != nil {
		return "", err
	}

	for _, s := range strObjects {
		hh.Write(unsafeGetBytes(s.String()))
	}

	b := make([]byte, 36)
	hex.Encode(b, hh.Sum(nil))
	for i, d := range dashes {
		value := b[d]
		b[d] = '-'
		b[len(b)-1-i] = value
	}

	return string(b), nil
}

// SetHashKeyWithSeed sets the seed and generates a new 32 key for the internal hash function
func SetHashKeyWithSeed(s int64) {
	_seed = rand.NewSource(s)
	r := rand.New(_seed)
	_, _ = r.Read(_key)
}

func unsafeGetBytes(s string) (b []byte) {
	if s == "" {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
