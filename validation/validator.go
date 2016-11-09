package validation

import (
	"strings"
)

type validator struct {
	data    map[string]interface{}
	rules   map[string][]string
	message string
}

func New(data map[string]interface{}, rules map[string]interface{}) *validator {
	return &validator{
		data:  data,
		rules: explodeRules(rules),
	}
}

func (v *validator) Fails() bool {
	return !v.passes()
}

func (v *validator) Passes() bool {
	for attribute, rules := range v.rules {
		for _, rule := range rules {
			if !v.validate(attribute, rule) {
				return false
			}
		}
	}

	return true
}

func (v *validator) GetMessage() {
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
