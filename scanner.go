package fingerprint

import (
	"encoding/hex"
	"fmt"
	"hash"
	"math/rand"
	"sync"
	"time"

	"github.com/minio/highwayhash"
)

// Scanner provides the ability to create a single hash.Hash that can be reused
type Scanner struct {
	hh hash.Hash
	mu sync.Mutex
}

// NewScanner returns an instance of scanner with the hash initialized using the math/rand package
func NewScanner() *Scanner {
	key := make([]byte, 32)
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	_, _ = r.Read(key)

	hh, _ := highwayhash.New128(key)
	return &Scanner{
		hh: hh,
	}
}

// Identify takes the string representation of objects and creates a hash in the form of a uuid
// if a seed is not provided a random 32 byte key is generated using a timestamp
func (sc *Scanner) Identify(strObjects ...fmt.Stringer) (string, error) {
	sc.mu.Lock()
	for _, s := range strObjects {
		sc.hh.Write(unsafeGetBytes(s.String()))
	}

	b := make([]byte, 36)
	hex.Encode(b, sc.hh.Sum(nil))
	sc.hh.Reset()
	sc.mu.Unlock()

	for i, d := range dashes {
		value := b[d]
		b[d] = '-'
		b[len(b)-1-i] = value
	}

	return string(b), nil
}

// SetHashKeyWithSeed sets the seed and generates a new 32 key for the internal hash function
func (sc *Scanner) SetHashKeyWithSeed(s int64) *Scanner {
	key := make([]byte, 32)
	source := rand.NewSource(s)
	r := rand.New(source)
	_, _ = r.Read(key)
	hh, _ := highwayhash.New128(key)

	sc.mu.Lock()
	sc.hh = hh
	sc.mu.Unlock()
	return sc
}
