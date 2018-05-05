package timezh

import (
	"testing"
	"time"

	"github.com/WindomZ/testify/assert"
)

func TestSince(t *testing.T) {
	assert.NotZero(t, Since(Now()))
	assert.True(t, Since(Now()) <= time.Millisecond)
}

func TestUntil(t *testing.T) {
	assert.NotZero(t, Until(Now()))
	assert.True(t, Until(Now()) <= time.Millisecond)
}

func TestUnix(t *testing.T) {
	assert.False(t, Unix(1, 0).IsZero())
}

func TestDate(t *testing.T) {
	assert.False(t, Date(1970, 1, 1, 0, 0, 0, 0, time.Local).IsZero())
}
