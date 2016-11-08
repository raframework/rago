package validation

import (
	"strings"
)

var ruleMethodMap = map[string]func(string, interface{}, []string) bool{
	"required": validateRequired,
}

type validator struct {
	data    map[string]interface{}
	rules   map[string][]string
	message string
}

func NewValidator(data interface{}, rules map[string]interface{}, messages messages) *validator {
	return &validator{
		data:  data,
		rules: explodeRules(rules),
	}
}

func (v *validator) fails() bool {
	return !v.passes()
}

func (v *validator) passes() bool {
	for attribute, rules := range v.rules {
		for _, rule := range rules {
			if !v.validate(attribute, rule) {
				return false
			}
		}
	}

	return true
}

func (v *validator) getMessage() {
	return v.message
}

func (v *validator) validate(attribute, rule string) bool {
	rule, parameters := parseRule(rule)
	value := v.getValue(attribute)

	if rule != "required" && value == nil {
		return true
	}

}

func (v *validator) getValue(attribute string) interface{} {
	value, ok := v.data[attribute]
	if !ok {
		return nil
	}

	return value
}

func parseRule(rule string) (string, []string) {
	parameters := []string{}
	if strings.Index(rule, ":") != -1 {
		splitRule := strings.SplitN(rule, ":", 2)
		rule = splitRule[0]
		parameters = parseParameters(splitRule[0], splitRule[1])
	}

	return rule, parameters
}

func explodeRules(rules map[string]interface{}) map[string][]string {
	r := map[string][]string{}

	for attribute, rule := range rules {
		switch rule.(type) {
		case string:
			r[attribute] = strings.Split(rule, "|")
		case []string:
			r[attribute] = rule
		default:
			panic("Invalid type of rule")
		}
	}

	return r
}

func parseParameters(rule, parameter string) []string {
	parameters := []string{}
	if rule == "regex" {
		parameters = append(parameters, parameter)
	} else {
		parameters = strings.Split(parameter, ",")
	}

	return parameters
}

func validateRequired(attribute string, value interface{}, parameters []string) bool {

}
