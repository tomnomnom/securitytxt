package securitytxt

import "testing"

func TestValidEmail(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"mail@example.com", true},
		{"invalid email", false},
	}

	for _, c := range cases {
		have := validEmail(c.in)
		if c.want != have {
			t.Errorf("want %t for validEmail(%s); have %t", c.want, c.in, have)
		}
	}
}

func TestValidURI(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"https://example.com", true},
		{"invalid URI", false},
	}

	for _, c := range cases {
		have := validURI(c.in)
		if c.want != have {
			t.Errorf("want %t for validURI(%s); have %t", c.want, c.in, have)
		}
	}
}

func TestValidPhone(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"+44 (5555) 555-555", true},
		{"invalid phone", false},
	}

	for _, c := range cases {
		have := validPhone(c.in)
		if c.want != have {
			t.Errorf("want %t for validPhone(%s); have %t", c.want, c.in, have)
		}
	}
}
