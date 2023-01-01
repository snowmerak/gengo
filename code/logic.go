package code

import "bytes"

func While(condition Code, body ...Code) Code {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("for " + string(condition) + " {\n")
	for _, line := range body {
		buf.WriteString("\t" + string(line) + "\n")
	}
	buf.WriteString("}\n")
	return Code(buf.String())
}

func For(init, condition, post Code, body ...Code) Code {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("for " + string(init) + "; " + string(condition) + "; " + string(post) + " {\n")
	for _, line := range body {
		buf.WriteString("\t" + string(line) + "\n")
	}
	buf.WriteString("}\n")
	return Code(buf.String())
}

func ForRange(key, value Code, collection Code, body ...Code) Code {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("for " + string(key) + ", " + string(value) + " := range " + string(collection) + " {\n")
	for _, line := range body {
		buf.WriteString("\t" + string(line) + "\n")
	}
	buf.WriteString("}\n")
	return Code(buf.String())
}

func Loop(body ...Code) Code {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("for {\n")
	for _, line := range body {
		buf.WriteString("\t" + string(line) + "\n")
	}
	buf.WriteString("}\n")
	return Code(buf.String())
}

func If(init Code, condition Code, body ...Code) Code {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("if ")
	if init != "" {
		buf.WriteString(string(init) + "; ")
	}
	buf.WriteString(string(condition) + " {\n")
	for _, line := range body {
		buf.WriteString("\t" + string(line) + "\n")
	}
	buf.WriteString("}\n")
	return Code(buf.String())
}

func IfElse(condition Code, ifBody []Code, elseBody []Code) Code {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("if " + string(condition) + " {\n")
	for _, line := range ifBody {
		buf.WriteString("\t" + string(line) + "\n")
	}
	buf.WriteString("} else {\n")
	for _, line := range elseBody {
		buf.WriteString("\t" + string(line) + "\n")
	}
	buf.WriteString("}\n")
	return Code(buf.String())
}

type SwitchCase struct {
	Condition Code
	Body      []Code
}

func Switch(value Code, cases []SwitchCase, defaultBody ...Code) Code {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("switch " + string(value) + " {\n")
	for _, c := range cases {
		buf.WriteString("case " + string(c.Condition) + ":\n")
		for _, line := range c.Body {
			buf.WriteString("\t" + string(line) + "\n")
		}
	}
	if len(defaultBody) > 0 {
		buf.WriteString("default:\n")
		for _, line := range defaultBody {
			buf.WriteString("\t" + string(line) + "\n")
		}
		buf.WriteString("}\n")
	}
	return Code(buf.String())
}

func Break() Code {
	return Code("break")
}

func Continue() Code {
	return Code("continue")
}

func Return(values ...Code) Code {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("return")
	for _, value := range values {
		buf.WriteString(" " + string(value))
	}
	return Code(buf.String())
}

func Goto(label string) Code {
	return Code("goto " + label)
}

func Label(label string) Code {
	return Code(label + ":")
}

func Select(cases []SwitchCase, defaultBody ...Code) Code {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("select {\n")
	for _, c := range cases {
		buf.WriteString("case " + string(c.Condition) + ":\n")
		for _, line := range c.Body {
			buf.WriteString("\t" + string(line) + "\n")
		}
	}
	if len(defaultBody) > 0 {
		buf.WriteString("default:\n")
		for _, line := range defaultBody {
			buf.WriteString("\t" + string(line) + "\n")
		}
		buf.WriteString("}\n")
	}
	return Code(buf.String())
}
