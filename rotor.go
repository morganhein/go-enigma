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
	Ring     int //the original ring offset due to the ring settings
	Offset   int //the current offset due to pressing keys
}

func makeRotors() []Rotor {
	return []Rotor{
		Rotor{ // Rotor I    - Royal
			Alphabet: "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
			Notches:  "R",
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

func (r *Rotor) rotate() bool {
	r.Offset++
	if r.Offset > 25 {
		r.Offset = 0
	}
	for _, notch := range strings.Split(r.Notches, "") {
		notchIndex := strings.Index(ALPHABET, notch)
		if r.Offset == notchIndex {
			return true
		}
	}
	return false
}

func (r *Rotor) getVisibleCharacter() string {
	return string(ALPHABET[r.Offset])
}

func (r *Rotor) getTranslation(letter string, d direction) string {
	initialLetter := ""
	defer func() {
		log.Debugw("rotor translation",
			"from", letter,
			"to", initialLetter,
			"alphabet", r.Alphabet)
	}()
	if d == LEFT {
		tOffset := strings.Index(ALPHABET, letter) + r.Offset
		initialLetter = r.getLetterFromAlphabet(tOffset, r.Alphabet)
		return initialLetter
	}
	//right
	tOffset := strings.Index(r.Alphabet, letter) - r.Offset
	initialLetter = r.getLetterFromAlphabet(tOffset, ALPHABET)
	return initialLetter
}

func (r *Rotor) getLetterFromAlphabet(index int, alphabet string) string {
	if index > 25 {
		index = index - 26
	}
	if index < 0 {
		index = 26 + index //index should be a negative number in this case
	}
	return string(alphabet[index])
}

//setRing is analogous to the Ringstellung
//http://practicalcryptography.com/ciphers/enigma-cipher/#the-ringstellung
func (r *Rotor) setRing(letter string) {
	r.Ring = strings.Index(ALPHABET, letter) - 1
}

func (r *Rotor) setRotor(letter string) {
	letter = strings.ToUpper(letter)
	r.Offset = int(letter[0]) - 65
}
