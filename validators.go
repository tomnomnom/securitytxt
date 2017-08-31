package securitytxt

import (
	"net/mail"
	"net/url"
	"regexp"
	"strings"
)

func validURI(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}

func validDisclosure(d string) bool {
	d = strings.ToLower(d)

	return d == "full" || d == "partial" || d == "none"
}

func validContact(c string) bool {
	_, err := mail.ParseAddress(c)
	validEmail := err != nil

	re := regexp.MustCompile("^\\+[0-9\\(\\) -]+$")
	validPhone := re.MatchString(c)

	return validEmail || validURI(c) || validPhone
}
