package code

func Declare(name string, typ Code) Code {
	return Code(name + " " + string(typ))
}

func Var(name string, typ Code, value Code) Code {
	t := Code("var " + name + " " + string(typ))
	if value != "" {
		t += " = " + value
	}
	return t
}

func ShortDeclare(name string, value Code) Code {
	return Code(name + " := " + string(value))
}

func Insert(name string, value Code) Code {
	return Code(name + " = " + string(value))
}
