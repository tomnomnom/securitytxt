package securitytxt

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const (
	contactField         = "contact"
	encryptionField      = "encryption"
	disclosureField      = "disclosure"
	acknowledgementField = "acknowledgement"
)

func parse(r io.Reader) (*File, error) {
	f := &File{}
	sc := bufio.NewScanner(r)

	n := 0
	for sc.Scan() {
		n++

		line := strings.TrimSpace(sc.Text())

		if line == "" {
			continue
		}

		if line[0] == '#' {
			f.addComment(line)
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			f.addError(fmt.Errorf("invalid input on line %d: %s", n, line))
			continue
		}

		option := strings.ToLower(parts[0])
		value := strings.TrimSpace(parts[1])

		switch option {

		case contactField:
			if !validContact(value) {
				f.addError(fmt.Errorf("invalid value '%s' for option '%s' on line %d", value, option, n))
				continue
			}
			f.addContact(value)

		case encryptionField:
			if !validURI(value) {
				f.addError(fmt.Errorf("invalid value '%s' for option '%s' on line %d", value, option, n))
				continue
			}
			f.addEncryption(value)

		case disclosureField:
			if !validDisclosure(value) {
				f.addError(fmt.Errorf(
					"invalid value '%s' for option '%s' on line %d, should be one of [full, partial, none]", value, option, n,
				))
				continue
			}
			f.addDisclosure(value)

		case acknowledgementField:
			if !validURI(value) {
				f.addError(fmt.Errorf("invalid value '%s' for option '%s' on line %d", value, option, n))
				continue
			}
			f.addAcknowledgement(value)

		default:
			f.addError(fmt.Errorf("invalid option '%s' on line %d", option, n))
		}

	}

	// No lines were scanned
	if n == 0 {
		f.addError(fmt.Errorf("empty file"))
	}

	if len(f.contact) < 1 {
		f.addError(fmt.Errorf("does not contain at least one contact field"))
	}

	if len(f.errors) > 0 {
		return f, fmt.Errorf("%d errors encountered during parsing", len(f.errors))
	}

	return f, nil
}
