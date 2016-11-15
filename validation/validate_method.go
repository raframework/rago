package validation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/raframework/rago/raerror"
	"github.com/raframework/rago/ralog"
)

const (
	REGEXP_EMAIL = "^[a-zA-Z0-9]+([_\\-.][a-zA-Z0-9]+)*@[a-zA-Z0-9]+([-.][a-zA-Z0-9]+)*\\.[a-zA-Z0-9]+([-.][a-zA-Z0-9]+)*$"
)

func requireParameterCount(count int, parameters []string, rule string) {
	if len(parameters) < count {
		raerror.PanicWith(raerror.TypInvalidArgument, 0, fmt.Sprintf("Validation rule %s requires at least %d parameters", rule, count))
	}
}

func getSize(value interface{}) float64 {
	switch value.(type) {
	case float64:
		return value.(float64)
	case string:
		return float64(len(value.(string)))
	case []string:
		return float64(len(value.([]string)))
	default:
		ralog.Critical("validation: invalid type of parameter for getSize(): ", value)
		raerror.PanicWith(raerror.TypInvalidArgument, 0, "validation: invalid type of parameter for getSize()")
	}

	return 0
}

func stringTofloat64(rule, s string) float64 {
	i, err := strconv.Atoi(s)
	if err != nil {
		raerror.PanicWith(
			raerror.TypInvalidArgument, 0, fmt.Sprintf("validation: invalid parameter for rule '%s', numeric of string is required", rule))
	}

	return float64(i)
}

func validateRequired(attribute string, value interface{}, parameters []string) bool {
	if value == nil {
		return false
	}

	switch value.(type) {
	case string:
		strValue, ok := value.(string)
		if !ok {
			return false
		}
		if strings.TrimSpace(strValue) == "" {
			return false
		}
	}

	return true
}

func validateString(attribute string, value interface{}, parameters []string) bool {
	ralog.Debug("validation: validateString value: ", value, " parameters: ", parameters)
	if value == nil {
		return false
	}

	switch value.(type) {
	case string:
		return true
	default:
		return false
	}
}

func validateBool(attribute string, value interface{}, parameters []string) bool {
	if value == nil {
		return false
	}

	switch value.(type) {
	case bool:
		return true
	default:
		return false
	}
}

func validateFloat(attribute string, value interface{}, parameters []string) bool {
	if value == nil {
		return false
	}

	switch value.(type) {
	case float64:
		return true
	default:
		return false
	}
}

func validateSize(attribute string, value interface{}, parameters []string) bool {
	return getSize(value) == stringTofloat64("size", parameters[0])
}

func validateMax(attribute string, value interface{}, parameters []string) bool {
	requireParameterCount(1, parameters, "max")

	return getSize(value) <= stringTofloat64("max", parameters[0])
}

func validateMin(attribute string, value interface{}, parameters []string) bool {
	requireParameterCount(1, parameters, "min")

	return getSize(value) >= stringTofloat64("min", parameters[0])
}

func validateBetween(attribute string, value interface{}, parameters []string) bool {
	requireParameterCount(2, parameters, "between")
	size := getSize(value)

	return size >= stringTofloat64("between", parameters[0]) && size <= stringTofloat64("between", parameters[1])
}

func validateEmail(attribute string, value interface{}, parameters []string) bool {
	strValue, ok := value.(string)
	if !ok {
		return false
	}
	if matched, _ := regexp.MatchString(REGEXP_EMAIL, strValue); !matched {
		return false
	}

	return true
}
