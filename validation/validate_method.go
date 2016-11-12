package validation

import (
	"regexp"
	"strings"
)

const (
	REGEXP_EMAIL = "^[a-zA-Z0-9]+([_\\-.][a-zA-Z0-9]+)*@[a-zA-Z0-9]+([-.][a-zA-Z0-9]+)*\\.[a-zA-Z0-9]+([-.][a-zA-Z0-9]+)*$"
)

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
