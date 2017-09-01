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

func (f *Field) setComments(cs []string) {
	f.comments = cs
}

// Comments returns the comments associated with a field
func (f Field) Comments() []string {
	return f.comments
}

func newField(option, value string) (*Field, error) {

	switch option {

	case contactField:
		if !validContact(value) {
			return nil, fmt.Errorf("invalid value '%s' for option '%s'", value, option)
		}

	case encryptionField:
		if !validURI(value) {
			return nil, fmt.Errorf("invalid value '%s' for option '%s'", value, option)
		}

	case disclosureField:
		if !validDisclosure(value) {
			return nil, fmt.Errorf("invalid value '%s' for option '%s'; should be one of [full, partial, none]", value, option)
		}

	case acknowledgementField:
		if !validURI(value) {
			return nil, fmt.Errorf("invalid value '%s' for option '%s'", value, option)
		}

	default:
		return nil, fmt.Errorf("invalid option '%s'", option)
	}

	return &Field{option: option, value: value}, nil
}
