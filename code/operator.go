package code

import "bytes"

func Add(a, b Code) Code {
	return Code(string(a) + " + " + string(b))
}

func AddAll(codes ...Code) Code {
	buf := bytes.NewBuffer(nil)
	for i, code := range codes {
		buf.WriteString(string(code))
		if i < len(codes)-1 {
			buf.WriteString(" + ")
		}
	}
	return Code(buf.String())
}

func Sub(a, b Code) Code {
	return Code(string(a) + " - " + string(b))
}

func SubAll(codes ...Code) Code {
	buf := bytes.NewBuffer(nil)
	for i, code := range codes {
		buf.WriteString(string(code))
		if i < len(codes)-1 {
			buf.WriteString(" - ")
		}
	}
	return Code(buf.String())
}

func Mul(a, b Code) Code {
	return Code(string(a) + " * " + string(b))
}

func MulAll(codes ...Code) Code {
	buf := bytes.NewBuffer(nil)
	for i, code := range codes {
		buf.WriteString(string(code))
		if i < len(codes)-1 {
			buf.WriteString(" * ")
		}
	}
	return Code(buf.String())
}

func Div(a, b Code) Code {
	return Code(string(a) + " / " + string(b))
}

func DivAll(codes ...Code) Code {
	buf := bytes.NewBuffer(nil)
	for i, code := range codes {
		buf.WriteString(string(code))
		if i < len(codes)-1 {
			buf.WriteString(" / ")
		}
	}
	return Code(buf.String())
}

func Mod(a, b Code) Code {
	return Code(string(a) + " % " + string(b))
}

func ModAll(codes ...Code) Code {
	buf := bytes.NewBuffer(nil)
	for i, code := range codes {
		buf.WriteString(string(code))
		if i < len(codes)-1 {
			buf.WriteString(" % ")
		}
	}
	return Code(buf.String())
}

func And(a, b Code) Code {
	return Code(string(a) + " && " + string(b))
}

func AndAll(codes ...Code) Code {
	buf := bytes.NewBuffer(nil)
	for i, code := range codes {
		buf.WriteString(string(code))
		if i < len(codes)-1 {
			buf.WriteString(" && ")
		}
	}
	return Code(buf.String())
}

func Or(a, b Code) Code {
	return Code(string(a) + " || " + string(b))
}

func OrAll(codes ...Code) Code {
	buf := bytes.NewBuffer(nil)
	for i, code := range codes {
		buf.WriteString(string(code))
		if i < len(codes)-1 {
			buf.WriteString(" || ")
		}
	}
	return Code(buf.String())
}

func Not(a Code) Code {
	return Code("!" + string(a))
}

func BitAnd(a, b Code) Code {
	return Code(string(a) + " & " + string(b))
}

func BitAndAll(codes ...Code) Code {
	buf := bytes.NewBuffer(nil)
	for i, code := range codes {
		buf.WriteString(string(code))
		if i < len(codes)-1 {
			buf.WriteString(" & ")
		}
	}
	return Code(buf.String())
}

func BitOr(a, b Code) Code {
	return Code(string(a) + " | " + string(b))
}

func BitOrAll(codes ...Code) Code {
	buf := bytes.NewBuffer(nil)
	for i, code := range codes {
		buf.WriteString(string(code))
		if i < len(codes)-1 {
			buf.WriteString(" | ")
		}
	}
	return Code(buf.String())
}

func BitClear(a, b Code) Code {
	return Code(string(a) + " &^ " + string(b))
}

func BitClearAll(codes ...Code) Code {
	buf := bytes.NewBuffer(nil)
	for i, code := range codes {
		buf.WriteString(string(code))
		if i < len(codes)-1 {
			buf.WriteString(" &^ ")
		}
	}
	return Code(buf.String())
}

func BitReverse(a Code) Code {
	return Code("^" + string(a))
}

func BitXor(a, b Code) Code {
	return Code(string(a) + " ^ " + string(b))
}

func BitXorAll(codes ...Code) Code {
	buf := bytes.NewBuffer(nil)
	for i, code := range codes {
		buf.WriteString(string(code))
		if i < len(codes)-1 {
			buf.WriteString(" ^ ")
		}
	}
	return Code(buf.String())
}

func BitLeftShift(a, b Code) Code {
	return Code(string(a) + " << " + string(b))
}

func BitRightShift(a, b Code) Code {
	return Code(string(a) + " >> " + string(b))
}
