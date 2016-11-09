package validation

var ruleMethodMap = map[string]func(string, interface{}, []string) bool{
	"required": validateRequired,
	"email":    validateEmail,
}

var defaultMessages = map[string]string{
	"required": "The :attribute field is required.",
	"email":    "The :attribute must be a valid email address.",
}
