package securitytxt

import "fmt"

// A Field is an option/value pair and its associated comments
type Field struct {
	option string
	value  string

	comments []string
}

// String returns the value of a field
func (f Field) String() string {
	return f.value
}

// Comments returns the comments associated with a field
func (f Field) Comments() []string {
	return f.comments
}

func newField(option, value string, comments []string) (Field, error) {

	switch option {

	case contactField:
		if !validContact(value) {
			return Field{}, fmt.Errorf("invalid value '%s' for option '%s'", value, option)
		}

	case encryptionField:
		if !validURI(value) {
			return Field{}, fmt.Errorf("invalid value '%s' for option '%s'", value, option)
		}

	case acknowledgmentsField:
		if !validURI(value) {
			return Field{}, fmt.Errorf("invalid value '%s' for option '%s'", value, option)
		}

	default:
		return Field{}, fmt.Errorf("invalid option '%s'", option)
	}

	return Field{option: option, value: value, comments: comments}, nil
}
