package validation

type ruleMethod func(string, interface{}, []string) bool

var ruleMethodMap = map[string]ruleMethod{
	"required": validateRequired,
	"email":    validateEmail,
}

var defaultMessages = map[string]string{
	"required": "The :attribute field is required.",
	"email":    "The :attribute must be a valid email address.",
}
