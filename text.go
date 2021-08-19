package faker

import "unicode/utf8"

type TextGen struct {
	Base string
	txtSize int
}

func (p *TextGen) GenText() string {
	if p.Base == "" {
		p.Base = "abcdefghijklmnopqrstuvwxyz"
	}

	builtTxt := ""
	for utf8.RuneCountInString(builtTxt) < p.txtSize {
		builtTxt += p.Base
	}
	builtTxt = string([]rune(builtTxt)[:p.txtSize])

	return builtTxt
}
