package go_enigma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeG(t *testing.T) {
	order := RotorOrder{
		LRotor: rotors[0],
		MRotor: rotors[1],
		RRotor: rotors[2],
	}
	output, err := Encode("g", "", "ZAA", order, ReflectorB)
	assert.NoError(t, err)
	assert.Equal(t, "P", output)
}

func TestRotorRightNotch(t *testing.T) {
	order := RotorOrder{
		LRotor: rotors[0],
		MRotor: rotors[1],
		RRotor: rotors[2],
	}
	output, err := Encode("A", "", "VAA", order, ReflectorB)
	assert.NoError(t, err)
	assert.Equal(t, "U", output)
}
func TestEncodeMessage(t *testing.T) {
	order := RotorOrder{
		LRotor: rotors[0],
		MRotor: rotors[1],
		RRotor: rotors[2],
	}
	//output, err := Encode("THIS IS A TEST MESSAGE DO YOU READ", "", "AAA", order, ReflectorB)
	output, err := Encode("T", "", "AAA", order, ReflectorB)
	assert.NoError(t, err)
	assert.Equal(t, "O", output)
	//assert.Equal(t, "OPGNDXCRWOMNLNECJZIJEDJGCXB", output)
}
