package code

import "bytes"

func Package(name string) Code {
	return Code("package " + name + "\n")
}

func Import(paths ...string) Code {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("import (\n")
	for _, path := range paths {
		buf.WriteString("\t\"" + path + "\"\n")
	}
	buf.WriteString(")\n")
	return Code(buf.String())
}
