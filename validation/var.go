package validation

type ruleMethod func(string, interface{}, []string) bool

var ruleMethodMap = map[string]ruleMethod{
	"required": validateRequired,

	"bool":   validateBool,
	"string": validateString,
	"float":  validateFloat,

	"size":    validateSize,
	"max":     validateMax,
	"min":     validateMin,
	"between": validateBetween,

	"email": validateEmail,
}

var defaultMessages = map[string]string{
	"required": "The :attribute field is required.",

	"int":    "The :attribute must be an integer.",
	"bool":   "The :attribute field must be true or false.",
	"string": "The :attribute must be a string.",
	"float":  "The :attribute must be a float.",

	"size":    "The size of :attribute must be :size.",
	"max":     "The size of :attribute may not be greater than :max.",
	"min":     "The size of :attribute must be at least :min.",
	"between": "The size of :attribute must be between :min and :max.",

	"email": "The :attribute must be a valid email address.",
}
