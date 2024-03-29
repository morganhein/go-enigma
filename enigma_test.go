package go_enigma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeMessage(t *testing.T) {
	order := RotorSettings{
		LRotor: rotors[0],
		MRotor: rotors[1],
		RRotor: rotors[2],
	}
	//output, err := Encode("THIS IS A TEST MESSAGE DO YOU READ", "", "AAA", order, ReflectorB)
	output, err := Encode("T", "", order, ReflectorB)
	assert.NoError(t, err)
	assert.Equal(t, "O", output)
	//assert.Equal(t, "OPGNDXCRWOMNLNECJZIJEDJGCXB", output)
}
