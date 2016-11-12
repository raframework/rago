package apperror

func PanicWith(typ int, code int) {
	panic(New(typ, code, ""))
}

func PanicWithMessage(typ int, code int, message string) {
	panic(New(typ, code, message))
}
