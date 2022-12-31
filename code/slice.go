package code

import "strings"

func Append(dest Code, source Code, elements ...Code) Code {
	buf := strings.Builder{}
	buf.WriteString(string(dest) + " = append(" + string(dest) + ", " + string(source))
	for _, element := range elements {
		buf.WriteString(", " + string(element))
	}
	buf.WriteString(")")
	return Code(buf.String())
}

func Slicing(slice Code, start Code, end Code) Code {
	buf := strings.Builder{}
	buf.WriteString(string(slice) + "[" + string(start) + ":" + string(end) + "]")
	return Code(buf.String())
}

func SlicingWithCap(slice Code, start Code, end Code, cap Code) Code {
	buf := strings.Builder{}
	buf.WriteString(string(slice) + "[" + string(start) + ":" + string(end) + ":" + string(cap) + "]")
	return Code(buf.String())
}
