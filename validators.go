package securitytxt

import (
	"net/mail"
	"net/url"
	"regexp"
)

func validURI(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}

func validEmail(e string) bool {
	_, err := mail.ParseAddress(e)
	return err == nil
}

func validPhone(p string) bool {
	re := regexp.MustCompile("^\\+[0-9\\(\\) -]+$")
	return re.MatchString(p)
}

func validContact(c string) bool {
	return validEmail(c) || validURI(c) || validPhone(c)
}
