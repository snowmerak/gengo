package main

import (
	"os"

	"github.com/snowmerak/gengo/code"
)

func main() {
	f := code.NewFile(os.Stdout)
	f.Write(
		code.Package("main"),
	).Write(
		code.Import("fmt"),
	).Write(
		code.Func(
			"main", nil, nil,
			[]code.Code{
				code.ShortDeclare("a", code.Array("int", code.Number(1), code.Number(2), code.Number(3))),
				code.ShortDeclare("r", code.Number(0)),
				code.ForRange("_", "v", "a",
					code.Insert("r", code.Add("r", "v")),
				),
				code.Call("fmt.Println", []code.Code{
					code.Code("r"),
				}),
			},
		),
	)
}
