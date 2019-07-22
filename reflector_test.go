package go_enigma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//just sanity check using some known values
func TestReflectorB(t *testing.T) {
	out := reflect("A", ReflectorB)
	assert.Equal(t, "Y", out)
	out = reflect("H", ReflectorB)
	assert.Equal(t, "D", out)
	out = reflect("Q", ReflectorB)
	assert.Equal(t, "E", out)
	out = reflect("Z", ReflectorB)
	assert.Equal(t, "T", out)
}
