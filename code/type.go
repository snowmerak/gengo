package code

import "strconv"

func TypeInt() Code {
	return Code("int")
}

func TypeInt8() Code {
	return Code("int8")
}

func TypeInt16() Code {
	return Code("int16")
}

func TypeInt32() Code {
	return Code("int32")
}

func TypeInt64() Code {
	return Code("int64")
}

func TypeUint() Code {
	return Code("uint")
}

func TypeUint8() Code {
	return Code("uint8")
}

func TypeUint16() Code {
	return Code("uint16")
}

func TypeUint32() Code {
	return Code("uint32")
}

func TypeUint64() Code {
	return Code("uint64")
}

func TypeUintptr() Code {
	return Code("uintptr")
}

func TypeFloat32() Code {
	return Code("float32")
}

func TypeFloat64() Code {
	return Code("float64")
}

func TypeComplex64() Code {
	return Code("complex64")
}

func TypeComplex128() Code {
	return Code("complex128")
}

func TypeBool() Code {
	return Code("bool")
}

func TypeString() Code {
	return Code("string")
}

func TypeByte() Code {
	return Code("byte")
}

func TypeRune() Code {
	return Code("rune")
}

func TypeError() Code {
	return Code("error")
}

func TypeInterface() Code {
	return Code("interface{}")
}

func TypeArray(length int, typ Code) Code {
	return Code("[" + strconv.Itoa(length) + "]" + string(typ))
}

func TypeSlice(typ Code) Code {
	return Code("[]" + string(typ))
}

func TypeMap(key, value Code) Code {
	return Code("map[" + string(key) + "]" + string(value))
}

func TypeSet(typ Code) Code {
	return Code("map[" + string(typ) + "]struct{}")
}

func TypeChan(typ Code) Code {
	return Code("chan " + string(typ))
}

func TypeChanSend(typ Code) Code {
	return Code("chan<- " + string(typ))
}

func TypeChanRecv(typ Code) Code {
	return Code("<-chan " + string(typ))
}

func TypeConvert(typ Code, value Code) Code {
	return Code(string(typ) + "(" + string(value) + ")")
}
