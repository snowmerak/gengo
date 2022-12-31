package code

import "bytes"

func Func(name string, params []Code, retuens []Code, body []Code) Code {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("func " + name + "(")
	for i, param := range params {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(string(param))
	}
	buf.WriteString(")")
	if len(retuens) > 0 {
		buf.WriteString(" ")
		for i, retuen := range retuens {
			if i > 0 {
				buf.WriteString(", ")
			}
			buf.WriteString(string(retuen))
		}
	}
	buf.WriteString(" {\n")
	for _, line := range body {
		buf.WriteString("\t" + string(line) + "\n")
	}
	buf.WriteString("}\n")
	return Code(buf.String())
}

func Call(name string, params []Code) Code {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(name + "(")
	for i, param := range params {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(string(param))
	}
	buf.WriteString(")")
	return Code(buf.String())
}
