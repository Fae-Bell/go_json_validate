package customjson

import "fmt"
import "strings"

type (
	TagExistsError struct {
		tagName string
	}
	ValidatorError interface {
		Error() string
	}
	ValidatorErrors struct {
		errs []ValidatorError
	}
	InvalidTypeError struct {
		givenType     string
		validTypes    []string
		validatorName string
	}
	ValidatorIdentifiers struct {
		fieldName string
		typeName  string
	}

	RequiredValidatorError struct {
		ValidatorIdentifiers
	}
	PatternMatchError struct {
		ValidatorIdentifiers
		pattern string
	}
)

func NewTagExistsError(tag string) *TagExistsError {
	return &TagExistsError{
		tagName: tag,
	}
}

func (e *TagExistsError) Error() string {
	return fmt.Sprintf("%s has already been registered", e.tagName)
}

func NewValidatorErrors() *ValidatorErrors {
	return &ValidatorErrors{}
}

func (e *ValidatorErrors) AppendErr(err ValidatorError) {
	if err != nil {
		e.errs = append(e.errs, err)
	}
}

func (e *ValidatorErrors) Error() string {
	s := make([]string, len(e.errs))
	for i, err := range e.errs {
		s[i] = err.Error()
	}
	return strings.Join(s, "\n")
}

func (e *ValidatorErrors) ContainsErrors() bool {
	return len(e.errs) != 0
}

func NewInvalidTypeError(givenType, validatorName string, validTypes []string) *InvalidTypeError {
	return &InvalidTypeError{
		givenType:     givenType,
		validatorName: validatorName,
		validTypes:    validTypes,
	}
}

func (e *InvalidTypeError) Error() string {
	return fmt.Sprintf(`invalid type "%s" given for validator "%s". only %v type(s) are accepted for this validator`, e.givenType, e.validatorName, e.validTypes)
}

func NewRequiredValidatorError(fieldName, typeName string) *RequiredValidatorError {
	return &RequiredValidatorError{
		ValidatorIdentifiers{
			fieldName: fieldName,
			typeName:  typeName,
		},
	}
}

func (e *RequiredValidatorError) Error() string {
	return fmt.Sprintf(`field "%s" on Type "%s" has been marked as required but has not been set`, e.fieldName, e.typeName)
}

func NewPatternMatchError(fieldName, typeName, pattern string) *PatternMatchError {
	return &PatternMatchError{
		ValidatorIdentifiers: ValidatorIdentifiers{
			fieldName: fieldName,
			typeName:  typeName,
		},
		pattern: pattern,
	}
}

func (e *PatternMatchError) Error() string {
	return fmt.Sprintf(`field "%s" on Type "%s" does not match the pattern provided of "%s"`, e.fieldName, e.typeName, e.pattern)
}
