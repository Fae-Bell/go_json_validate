package customjson

import (
	"fmt"
)

type (
	TagRegister func(typeName, fieldName, tagValue string, jsonValue interface{}) error
)

var registry = map[string]TagRegister{}

func init() {
	registry["json.required"] = valueRequired
}

func valueRequired(typeName, fieldName, tagValue string, jsonValue interface{}) error {
	if jsonValue == nil && tagValue != "false" {
		return fmt.Errorf(`Field "%s" on Type "%s" has been marked as required but has not been set`, fieldName, typeName)
	}
	return nil
}
