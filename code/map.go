package code

func MapInsert(m Code, k Code, v Code) Code {
	return Code(string(m) + "[" + string(k) + "] = " + string(v))
}

func MapDelete(m Code, k Code) Code {
	return Code("delete(" + string(m) + ", " + string(k) + ")")
}

func MapLookupDeclare(v Code, ok Code, m Code, k Code) Code {
	return Code(string(v) + ", " + string(ok) + " := " + string(m) + "[" + string(k) + "]")
}

func MapLookup(v Code, ok Code, m Code, k Code) Code {
	return Code(string(v) + " " + string(ok) + " = " + string(m) + "[" + string(k) + "]")
}
