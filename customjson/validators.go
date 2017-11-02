package customjson

import (
	"fmt"
	"regexp"
)

const (
	PatternValidatorName = "pattern"
)

func valueRequired(typeName, fieldName, tagValue string, jsonValue interface{}) ValidatorError {
	if jsonValue == nil && tagValue != "false" {
		return NewRequiredValidatorError(fieldName, typeName)
	}
	return nil
}

func valuePatternMatch(typeName, fieldName, tagValue string, jsonValue interface{}) ValidatorError {
	if jsonValue == nil && tagValue != "" {
		return NewPatternMatchError(fieldName, typeName, tagValue)
	}
	var sValue string
	switch jsonValue.(type) {
	case string:
		sValue = jsonValue.(string)
	case *string:
		sValue = *jsonValue.(*string)
	default:
		return NewInvalidTypeError(fmt.Sprintf("%T", jsonValue), PatternValidatorName, []string{"string", "*string"})
	}

	if m, _ := regexp.MatchString(tagValue, sValue); !m {
		return NewPatternMatchError(fieldName, typeName, tagValue)
	}
	return nil
}
