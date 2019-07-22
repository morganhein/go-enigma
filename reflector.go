package go_enigma

import (
	"strings"

	"golang.org/x/xerrors"
)

type Reflector string

func (r Reflector) validate() error {
	for k, v := range r {
		index := strings.Index(ALPHABET, string(v))
		if r[index] != ALPHABET[k] {
			return xerrors.New("reflector is malformed")
		}
	}
	return nil
}
