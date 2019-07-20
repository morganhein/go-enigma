package go_enigma

import "strings"

//https://charlesreid1.github.io/enigma-cipher-implementation-part-3-enigma-in-java-without-objects.html
/*
define plaintext message
define normal alphabet and scrambled alphabets
define list of switchboard swap pairs
define list of reflector swap pairs
for each character in plaintext message:

    # Apply switchboard transformation
    for each pair in switchboard swap pairs:
        if character in swap pair, swap its value

    # Apply forward rotor transformation
    for each scrambled alphabet:
        get index of character in normal alphabet
        get new character at that index in scrambled alphabet
        replace character with new character

    # Apply reflector transformation
    for each pair in reflector swap pairs:
        if character in swap pair, swap its value

    # Apply reverse rotor transformation
    for each scrambled alphabet:
        get index of input character in scrambled alphabet
        get new character at that index in normal alphabet
        replace character with new character

    # Apply switchboard transformation
    for each pair in switchboard swap pairs:
        if character in swap pair, swap its value

    concatenate transformed input character to ciphertext message

    # Increment rotor wheels
    for each rotor/scrambled alphabet, left to right:
        get index of left notch in left alphabet
        get index of right notch in right alphabet
        if left index equals right index:
            cycle left alphabet forward 1 character
    cycle right-most alphabet forward 1 character
*/
type Rotor struct {
	Layout string
	Notch  string
}

const (
	ALPHABET string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	rotors = []Rotor{
		Rotor{ // Rotor I    - Royal
			Layout: "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
			Notch:  "R",
		},
		Rotor{ // Rotor II   - Flags
			Layout: "AJDKSIRUXBLHWTMCQGZNPYFVOE",
			Notch:  "F",
		},
		Rotor{ // Rotor V - Wave
			Layout: "BDFHJLCPRTXVZNYEIWGAKMUSQO",
			Notch:  "W",
		},
		Rotor{ // Rotor IV   - Kings
			Layout: "ESOVPZJAYQUIRHXLNFTGKDCMWB",
			Notch:  "K",
		},
		Rotor{ // Rotor V   - Above
			Layout: "VZBRGITYUPSDNHLXAWMJQOFECK",
			Notch:  "A",
		},
		Rotor{
			Layout: "JPGVOUMFYQBENHZRDKASXLICTW",
			Notch:  "AN",
		},
		Rotor{
			Layout: "NZJHGRCXMYSWBOUFAIVLPEKQDT",
			Notch:  "AN",
		},
		Rotor{
			Layout: "FKQHTLXOCBJSPDZRAMEWNIUYGV",
			Notch:  "AN",
		},
	}
)

type RotorOrder struct {
	LRotor Rotor
	MRotor Rotor
	RRotor Rotor
}

func Encode(message string, switchBoardPairs string, rotors RotorOrder) string {
	workingStr := ""
	//perform switchboard swapping
	for _, v := range strings.Split(message, "") {
		workingStr += swapPairs(v, switchBoardPairs)
	}

	//rotate the rotors first

	//Right to Left through rotors

	//reflector substitution

	return ""
}

func swapPairs(letter string, switchBoardPairs string) string {
	for _, v := range strings.Split(switchBoardPairs, ":") {
		if len(v) != 2 {
			continue
		}
		if letter == string(v[0]) {
			return string(v[1])
		}
		if letter == string(v[1]) {
			return string(v[1])
		}
	}
	return letter
}

func passThroughRotor(letter string, rotor Rotor, direction string)
