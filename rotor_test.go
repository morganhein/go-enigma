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
	l = r.getTranslation()
}

func TestRotate(t *testing.T) {
	r := Rotor{
		Offset: 0,
	}
	r.rotate()
	assert.Equal(t, 1, r.Offset)
	r.rotate()
	assert.Equal(t, 2, r.Offset)
	r.Offset = 2
	r.rotate()
	assert.Equal(t, 3, r.Offset)
}
