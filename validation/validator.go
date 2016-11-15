package validation

import (
	"errors"
	"strings"

	"github.com/raframework/rago/raerror"
	"github.com/raframework/rago/ralog"
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

func explodeRules(rules map[string]interface{}) map[string][]string {
	r := map[string][]string{}

	for attribute, rule := range rules {
		switch rule.(type) {
		case string:
			strRule, _ := rule.(string)
			r[attribute] = strings.Split(strRule, "|")
		case []string:
			sliceRule, _ := rule.([]string)
			r[attribute] = sliceRule
		default:
			panic("Invalid type of rule")
		}
	}

	return r
}

func (v *validator) Fails() bool {
	return !v.Passes()
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

func (v *validator) GetMessage() string {
	return v.message
}

func (v *validator) validate(attribute, rule string) bool {
	ralog.Debug("validation: validate, attribute: ", attribute, " rule: ", rule)
	rule, parameters := parseRule(rule)
	value := v.getValue(attribute)

	if rule != "required" && value == nil {
		return true
	}

	method, err := getRuleMethod(rule)
	ralog.Debug("validation: method: ", method)
	if err != nil {
		e := "validation: rule " + rule + " not supported"
		ralog.Critical(e)
		raerror.PanicWith(raerror.TypInvalidArgument, 0, e)
	}

	// Call the rule method
	if !method(attribute, value, parameters) {
		message, _ := defaultMessages[rule]
		if rule == "size" {
			message = strings.Replace(message, ":size", parameters[0], -1)
		} else if rule == "max" {
			message = strings.Replace(message, ":max", parameters[0], -1)
		} else if rule == "min" {
			message = strings.Replace(message, ":min", parameters[0], -1)
		} else if rule == "between" {
			message = strings.Replace(message, ":min", parameters[0], -1)
			message = strings.Replace(message, ":max", parameters[1], -1)
		}
		message = strings.Replace(message, ":attribute", attribute, -1)
		v.message = message

		return false
	}

	return true
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

func (v *validator) getValue(attribute string) interface{} {
	value, ok := v.data[attribute]
	if !ok {
		return nil
	}

	return value
}

func getRuleMethod(rule string) (ruleMethod, error) {
	method, ok := ruleMethodMap[rule]
	if !ok {
		return nil, errors.New("rule '" + rule + "'method not found")
	}

	return method, nil
}

func parseParameters(rule, parameter string) []string {
	parameters := []string{}
	if rule == "regex" {
		parameters = append(parameters, parameter)
	} else {
		parameters = strings.Split(parameter, ",")
	}

	for key, value := range parameters {
		parameters[key] = strings.TrimSpace(value)
	}

	return parameters
}
