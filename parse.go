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

	commentBuffer := make([]string, 0)
	n := 0
	for sc.Scan() {
		n++

		line := strings.TrimSpace(sc.Text())

		if line == "" {
			// Reset the comment buffer because the comments
			// we've stored aren't associated with any field
			commentBuffer = commentBuffer[:0]
			continue
		}

		if line[0] == '#' {
			f.addComment(line)
			commentBuffer = append(commentBuffer, line)
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			f.addError(fmt.Errorf("invalid input on line %d: %s", n, line))
			commentBuffer = commentBuffer[:0]
			continue
		}

		option := strings.ToLower(parts[0])
		value := strings.TrimSpace(parts[1])

		field, err := newField(option, value, commentBuffer)
		commentBuffer = commentBuffer[:0]

		if err != nil {
			f.addError(fmt.Errorf("%s on line %d", err, n))
			continue
		}

		f.addField(field)

	}

	// No lines were scanned
	if n == 0 {
		f.addError(fmt.Errorf("empty file"))
	}

	if len(f.Contact()) < 1 {
		f.addError(fmt.Errorf("does not contain at least one contact field"))
	}

	if len(f.Errors()) > 0 {
		return f, fmt.Errorf("%d errors encountered during parsing", len(f.errors))
	}

	return f, nil
}
