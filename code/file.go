package code

import "io"

type Code string

type File struct {
	writer io.Writer
}

func NewFile(writer io.Writer) *File {
	return &File{writer}
}

func (f *File) Write(code Code) *File {
	f.writer.Write([]byte(code))
	return f
}
