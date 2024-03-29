package go_enigma

import (
	"strings"

	"go.uber.org/zap"
)

type direction bool

const (
	ALPHABET   string    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ReflectorA Reflector = "EJMZALYXVBWFCRQUONTSPIKHGD"
	ReflectorB Reflector = "YRUHQSLDPXNGOKMIEBFZCWVJAT"
	ReflectorC Reflector = "FVPJIAOYEDRZXWGCTKUQSBNMHL"
	LEFT       direction = true
	RIGHT      direction = false
)

type RotorSettings struct {
	LRotor Rotor // Left
	MRotor Rotor // Middle
	RRotor Rotor // Right
}

var (
	log *zap.SugaredLogger
)

func init() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	log = logger.Sugar()
}

func Encode(message, plugBoardPairs string, rotors RotorSettings, reflector Reflector) (string, error) {
	if err := reflector.validate(); err != nil {
		return "", err
	}
	//convert all characters to uppercase, since that's the ASCII range all operations are done in
	//and strip all spaces
	message = strings.Replace(strings.ToUpper(message), " ", "", -1)
	log.Infof("Encoding message: %v", message)

	log.Debugw("Input parameters",
		"plugBoardPairs", plugBoardPairs,
		"rotors", rotors,
		"reflector", reflector)

	workingStr := ""
	//switchboard swapping
	for _, v := range strings.Split(message, "") {
		workingStr += swapPairs(v, plugBoardPairs)
	}
	log.Infow("After initial switchboard swapping", "message", workingStr)

	rotorString := ""
	//for each character in the message:
	for _, v := range strings.Split(workingStr, "") {
		log.Debugf("input character: %v", v)
		//rotate the rotors first
		rotors = rotate(rotors)
		log.Info("Grundstellung: ", string(rotors.LRotor.Offset+65)+string(rotors.MRotor.Offset+65)+
			string(rotors.RRotor.Offset+65))
		//Right to Left through rotors
		v = translate(rotors, LEFT, v)
		log.Debugf("after right to left: %v", v)
		//reflector substitution
		v = reflect(v, reflector)
		log.Debugf("after reflection: %v", v)
		//Left to Right through rotors
		v = translate(rotors, RIGHT, v)
		log.Debugf("after left to right: %v", v)

		rotorString += v
	}

	workingStr = ""
	//switchboard swapping
	for _, v := range strings.Split(rotorString, "") {
		workingStr += swapPairs(v, plugBoardPairs)
	}
	log.Infof("after post switchboard swapping: %v\n", workingStr)
	return workingStr, nil
}

func swapPairs(letter string, plugBoardPairs string) string {
	for _, v := range strings.Split(plugBoardPairs, ":") {
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

func rotate(rotors RotorSettings) RotorSettings {
	log.Debug("rotating rotors")
	rnotch := rotors.RRotor.rotate()
	mnotch := false
	if rnotch {
		mnotch = rotors.MRotor.rotate()
	}
	if mnotch {
		rotors.LRotor.rotate()
	}
	return rotors
}

func translate(rotors RotorSettings, d direction, letter string) string {
	//if we're translating from right to left
	if d == LEFT {
		letter = rotors.RRotor.getTranslation(letter, LEFT)
		letter = rotors.MRotor.getTranslation(letter, LEFT)
		letter = rotors.LRotor.getTranslation(letter, LEFT)
		return letter
	}
	//if we're translating left to right
	letter = rotors.LRotor.getTranslation(letter, RIGHT)
	letter = rotors.MRotor.getTranslation(letter, RIGHT)
	letter = rotors.RRotor.getTranslation(letter, RIGHT)
	return letter
}

func reflect(letter string, r Reflector) string {
	index := strings.Index(ALPHABET, letter)
	return string(r[index])
}
