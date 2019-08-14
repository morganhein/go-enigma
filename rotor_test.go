package go_enigma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//test going left through a rotor
func TestRotorRingAlignmentLeft(t *testing.T) {
	r := makeRotors()[0]
	r.Offset = 0
	r.Ring = 0
	letter := r.getTranslation("A", LEFT)
	assert.Equal(t, "E", letter)
}

//test going right through a rotor
func TestRotorRingAlignmentRight(t *testing.T) {
	r := makeRotors()[0]
	r.Offset = 0
	r.Ring = 0
	letter := r.getTranslation("Z", RIGHT)
	assert.Equal(t, "J", letter)
}

func TestGetTranslation(t *testing.T) {
	//follows https://www.codesandciphers.org.uk/enigma/example1.htm
	r := Rotor{
		Alphabet: "BDFHJLCPRTXVZNYEIWGAKMUSQO",
		Notches:  "R",
	}
	result := r.getTranslation("G", LEFT)
	assert.Equal(t, "C", result)

	m := Rotor{
		Alphabet: "AJDKSIRUXBLHWTMCQGZNPYFVOE",
		Notches:  "R",
		Offset:   0,
	}
	result = m.getTranslation(result, LEFT)
	assert.Equal(t, "D", result)

	l := Rotor{
		Alphabet: "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		Notches:  "R",
		Offset:   0,
	}
	result = l.getTranslation(result, LEFT)
	assert.Equal(t, "F", result)

	result = l.getTranslation("S", RIGHT)
	assert.Equal(t, "S", result)

	result = m.getTranslation(result, RIGHT)
	assert.Equal(t, "E", result)

	result = r.getTranslation(result, RIGHT)
	assert.Equal(t, "P", result)
}

func TestGetTranslationWithOffset(t *testing.T) {
	r := Rotor{
		Alphabet: "BDFHJLCPRTXVZNYEIWGAKMUSQO",
		Notches:  "R",
	}
	r.setRotor("B")
	l := r.getTranslation("A", LEFT)
	assert.Equal(t, "D", l)

	r.setRotor("Z")
	l = r.getTranslation("A", LEFT)
	assert.Equal(t, "O", l)

	r = makeRotors()[0]
	r.setRotor("B")

	r.setRotor("A")
	l = r.getTranslation("A", RIGHT)
	assert.Equal(t, "U", l)

	r.setRotor("Z")
	l = r.getTranslation("F", RIGHT)
	assert.Equal(t, "U", l)
}

func TestRotate(t *testing.T) {
	r := Rotor{}
	r.rotate()
	assert.Equal(t, 1, r.Offset)
	r.rotate()
	assert.Equal(t, 2, r.Offset)
	r.rotate()
	assert.Equal(t, 3, r.Offset)
}

func TestGetLetterFromAlphabet(t *testing.T) {
	r := Rotor{}
	l := r.getLetterFromAlphabet(0, ALPHABET)
	assert.Equal(t, "A", l)

	l = r.getLetterFromAlphabet(1, ALPHABET)
	assert.Equal(t, "B", l)

	l = r.getLetterFromAlphabet(-1, ALPHABET)
	assert.Equal(t, "Z", l)

	l = r.getLetterFromAlphabet(-2, ALPHABET)
	assert.Equal(t, "Y", l)

	l = r.getLetterFromAlphabet(26, ALPHABET)
	assert.Equal(t, "A", l)
}

func TestCodesAndCiphersExample(t *testing.T) {
	//https://www.codesandciphers.org.uk/enigma/example1.htm
	rotors := makeRotors()
	rs := RotorSettings{
		LRotor: rotors[0],
		MRotor: rotors[1],
		RRotor: rotors[2],
	}
	rs.RRotor.setRotor("Z")
	result, _ := Encode("G", "", rs, ReflectorB)
	assert.Equal(t, "P", result)
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

func TestBOffset(t *testing.T) {
	rotors := makeRotors()
	order := RotorSettings{
		LRotor: rotors[0],
		MRotor: rotors[1],
		RRotor: rotors[2],
	}
	order.RRotor.setRotor("B")
	output, err := Encode("A", "", order, ReflectorB)
	assert.NoError(t, err)
	assert.Equal(t, "D", output)
}
