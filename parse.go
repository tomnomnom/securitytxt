package securitytxt

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

const (
	contactField         = "contact"
	encryptionField      = "encryption"
	disclosureField      = "disclosure"
	acknowledgementField = "acknowledgement"
)

type File struct {
	errors   []error
	comments []string

	contact         []string
	encryption      []string
	disclosure      []string
	acknowledgement []string
}

func (f *File) parse(r io.Reader) error {
	sc := bufio.NewScanner(r)

	n := 0
	for sc.Scan() {
		n++

		line := strings.TrimSpace(sc.Text())

		if line == "" {
			continue
		}

		if line[0] == '#' {
			f.AddComment(line)
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
				log.Println(value)
				continue
			}
			f.AddContact(value)

		case encryptionField:
			if !validURI(value) {
				f.addError(fmt.Errorf("invalid value '%s' for option '%s' on line %d", value, option, n))
				continue
			}
			f.AddEncryption(value)

		case disclosureField:
			if !validDisclosure(value) {
				f.addError(fmt.Errorf(
					"invalid value '%s' for option '%s' on line %d, should be one of [full, partial, none]", value, option, n,
				))
				continue
			}
			f.AddDisclosure(value)

		case acknowledgementField:
			if !validURI(value) {
				f.addError(fmt.Errorf("invalid value '%s' for option '%s' on line %d", value, option, n))
				continue
			}
			f.AddAcknowledgement(value)

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
		return fmt.Errorf("%d errors encountered during parsing", len(f.errors))
	}

	return nil
}

func (f *File) AddContact(c string) {
	f.contact = append(f.contact, c)
}

func (f File) Contact() []string {
	return f.contact
}

func (f *File) AddEncryption(e string) {
	f.encryption = append(f.encryption, e)
}

func (f File) Encryption() []string {
	return f.encryption
}

func (f *File) AddDisclosure(d string) {
	f.disclosure = append(f.disclosure, d)
}

func (f File) Disclosure() []string {
	return f.disclosure
}

func (f *File) AddAcknowledgement(a string) {
	f.acknowledgement = append(f.acknowledgement, a)
}

func (f File) Acknowledgement() []string {
	return f.acknowledgement
}

func (f *File) AddComment(c string) {
	f.comments = append(f.comments, c)
}

func (f File) Comments() []string {
	return f.comments
}

func (f *File) addError(err error) {
	f.errors = append(f.errors, err)
}

func (f File) Errors() []error {
	return f.errors
}
