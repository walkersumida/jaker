package text

import "unicode/utf8"

type Text struct {
	Base string
	Size int
}

func New() *Text {
	return &Text{}
}

func (p *Text) Gen() string {
	if p.Base == "" {
		p.Base = "abcdefghijklmnopqrstuvwxyz"
	}

	builtTxt := ""
	for utf8.RuneCountInString(builtTxt) < p.Size {
		builtTxt += p.Base
	}
	builtTxt = string([]rune(builtTxt)[:p.Size])

	return builtTxt
}
