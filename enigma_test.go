package go_enigma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeG(t *testing.T) {
	rotors := makeRotors()
	rs := RotorSettings{
		LRotor: rotors[0],
		MRotor: rotors[1],
		RRotor: rotors[2],
	}
	rs.RRotor.Offset = 25
	output, err := Encode("g", "", rs, ReflectorB)
	assert.NoError(t, err)
	assert.Equal(t, "P", output)
}

func TestRotorRightNotch(t *testing.T) {
	rotors := makeRotors()
	order := RotorSettings{
		LRotor: rotors[0],
		MRotor: rotors[1],
		RRotor: rotors[2],
	}
	order.RRotor.setRotor("V")

	output, err := Encode("A", "", order, ReflectorB)
	assert.NoError(t, err)
	assert.Equal(t, "U", output)
}

//func TestEncodeMessage(t *testing.T) {
//	order := RotorSettings{
//		LRotor: rotors[0],
//		MRotor: rotors[1],
//		RRotor: rotors[2],
//	}
//	//output, err := Encode("THIS IS A TEST MESSAGE DO YOU READ", "", "AAA", order, ReflectorB)
//	output, err := Encode("T", "", "AAA", order, ReflectorB)
//	assert.NoError(t, err)
//	assert.Equal(t, "O", output)
//	//assert.Equal(t, "OPGNDXCRWOMNLNECJZIJEDJGCXB", output)
//}
