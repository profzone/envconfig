package envconfig

import "strings"

type Password string

func (p Password) String() string {
	return string(p)
}

func (p Password) SecurityString() string {
	var r []rune
	if len(p) > 10 {
		return strings.Repeat("*", 10)
	} else {
		for range []rune(string(p)) {
			r = append(r, []rune("*")...)
		}
		return string(r)
	}
}
