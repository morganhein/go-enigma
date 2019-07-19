package go_enigma

//https://charlesreid1.github.io/enigma-cipher-implementation-part-3-enigma-in-java-without-objects.html

const ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var (
	wheel = []string{
		ALPHABET,
		"EKMFLGDQVZNTOWYHXUSPAIBRCJ", // Rotor I    - Royal
		"AJDKSIRUXBLHWTMCQGZNPYFVOE", // Rotor II   - Flags
		"BDFHJLCPRTXVZNYEIWGAKMUSQO", // Rotor III  - Wave
		"ESOVPZJAYQUIRHXLNFTGKDCMWB", // Rotor IV   - Kings
		"VZBRGITYUPSDNHLXAWMJQOFECK", // Rotor V    - Above
		//"JPGVOUMFYQBENHZRDKASXLICTW",
		//"NZJHGRCXMYSWBOUFAIVLPEKQDT",
		//"FKQHTLXOCBJSPDZRAMEWNIUYGV",
	}
	notch = []string{
		"",  // No notch
		"R", // Royal
		"F", // Flags
		"W", // Wave
		"K", // Kings
		"A", // Above
	}
)
