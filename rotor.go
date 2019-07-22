package go_enigma

import (
	"strings"
)

var (
	rotors = makeRotors()
)

type Rotor struct {
	Alphabet string
	Notches  string
	Ring     int
	Offset   int
}

func makeRotors() []Rotor {
	return []Rotor{
		Rotor{ // Rotor I    - Royal
			Alphabet: "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
			Notches:  "R",
			Ring:     1,
		},
		Rotor{ // Rotor II   - Flags
			Alphabet: "AJDKSIRUXBLHWTMCQGZNPYFVOE",
			Notches:  "F",
		},
		Rotor{ // Rotor III - Wave
			Alphabet: "BDFHJLCPRTXVZNYEIWGAKMUSQO",
			Notches:  "W",
		},
		Rotor{ // Rotor IV   - Kings
			Alphabet: "ESOVPZJAYQUIRHXLNFTGKDCMWB",
			Notches:  "K",
		},
		Rotor{ // Rotor V   - Above
			Alphabet: "VZBRGITYUPSDNHLXAWMJQOFECK",
			Notches:  "ReflectorA",
		},
		Rotor{
			Alphabet: "JPGVOUMFYQBENHZRDKASXLICTW",
			Notches:  "AN",
		},
		Rotor{
			Alphabet: "NZJHGRCXMYSWBOUFAIVLPEKQDT",
			Notches:  "AN",
		},
		Rotor{
			Alphabet: "FKQHTLXOCBJSPDZRAMEWNIUYGV",
			Notches:  "AN",
		},
	}
}

func (r *Rotor) rotate() {
	b := r.current[0]
	b++
	if b == 91 {
		b = 65
	}
	r.current = string(b)
}

func (r *Rotor) getTranslation(letter string, d direction) string {

}

//setAlignment is analogous to the Ringstellung
//http://practicalcryptography.com/ciphers/enigma-cipher/#the-ringstellung
func (r *Rotor) setShift(letter string) {
	r.ring = strings.Index(ALPHABET, letter)
}
