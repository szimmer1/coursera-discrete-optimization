package fib

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"time"
)

// i = [0, 1, 2, 3, 4, 5,  6,  ...]
// n = [1, 2, 3, 5, 8, 13, 21, ...]
func TestFib(t *testing.T) {
	assert.Equal(t, 1, Fib(0, nil))
	assert.Equal(t, 2, Fib(1, nil))
	assert.Equal(t, 3, Fib(2, nil))
	assert.Equal(t, 21, Fib(6, nil))

	// ~1e13 checks
	i := int(math.Log2(float64(1e13)))
	start := time.Now()
	res := Fib(i, nil)
	d := time.Since(start)
	t.Logf("i=%d: %d, %s\n", i, res, d)
}

func TestFibMemo(t *testing.T) {
	assert.Equal(t, 1, FibMemo(0))
	assert.Equal(t, 2, FibMemo(1))
	assert.Equal(t, 3, FibMemo(2))
	assert.Equal(t, 21, FibMemo(6))

	i := int(math.Log2(float64(1e13)))
	start := time.Now()
	res := FibMemo(i)
	d := time.Since(start)
	t.Logf("i=%d: %d, %s\n", i, res, d)
}
