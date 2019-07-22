package go_enigma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotorRingAlignmentLeft(t *testing.T) {
	r := rotors[0]
	r.current = "A"
	r.setShift("B")
	letter := r.getTranslation("A", LEFT)
	assert.Equal(t, "K", letter)
}

func TestRotorRingAlignmentRight(t *testing.T) {
	r := rotors[0]
	r.current = "A"
	r.setShift("B")
	letter := r.getTranslation("A", RIGHT)
	assert.Equal(t, "T", letter)
}

func TestGetTranslation(t *testing.T) {
	//follows https://www.codesandciphers.org.uk/enigma/example1.htm
	r := Rotor{
		Layout:  "BDFHJLCPRTXVZNYEIWGAKMUSQO",
		Notch:   "R",
		current: "A",
	}
	result := r.getTranslation("G", LEFT)
	assert.Equal(t, "C", result)

	m := Rotor{
		Layout:  "AJDKSIRUXBLHWTMCQGZNPYFVOE",
		Notch:   "R",
		current: "A",
	}
	result = m.getTranslation(result, LEFT)
	assert.Equal(t, "D", result)

	l := Rotor{
		Layout:  "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		Notch:   "R",
		current: "A",
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

func TestRotate(t *testing.T) {
	r := Rotor{
		current: "A",
	}
	r.rotate()
	assert.Equal(t, "B", r.current)
	r.rotate()
	assert.Equal(t, "C", r.current)
	r.current = "Z"
	r.rotate()
	assert.Equal(t, "A", r.current)
}
