package code

import (
	"bytes"
	"strconv"
)

func Number(n int) Code {
	return Code(strconv.Itoa(n))
}

func String(s string) Code {
	return Code(strconv.Quote(s))
}

func Bool(b bool) Code {
	return Code(strconv.FormatBool(b))
}

func Float(f float64) Code {
	return Code(strconv.FormatFloat(f, 'f', -1, 64))
}

func Byte(b byte) Code {
	return Code(strconv.QuoteRune(rune(b)))
}

func Rune(r rune) Code {
	return Code(strconv.QuoteRune(r))
}

func Array(typ Code, values ...Code) Code {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("[]")
	buf.WriteString(string(typ))
	buf.WriteString("{")
	for i, value := range values {
		buf.WriteString(string(value))
		if i < len(values)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteString("}")
	return Code(buf.String())
}
