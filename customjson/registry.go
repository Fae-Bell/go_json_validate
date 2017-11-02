package customjson

type (
	TagRegister func(typeName, fieldName, tagValue string, jsonValue interface{}) ValidatorError
)

var registry = map[string]TagRegister{}

func init() {
	registry["json.required"] = valueRequired
	registry["json.pattern"] = valuePatternMatch
}

func RegisterTagValidator(name string, validator TagRegister, overwrite bool) error {
	if _, exists := registry[name]; exists && !overwrite {
		return NewTagExistsError(name)
	}
	registry[name] = validator
	return nil
}
