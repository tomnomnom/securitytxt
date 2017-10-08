package securitytxt

import "testing"

func TestFilters(t *testing.T) {
	f := &File{}
	f.addField(Field{option: contactField, value: "https://example.com"})
	f.addField(Field{option: contactField, value: "+44 5555 555 555"})
	f.addField(Field{option: contactField, value: "mail@example.com"})

	if len(f.PhoneContact()) != 1 {
		t.Errorf("want 1 phone contact; have %d", len(f.PhoneContact()))
	}

	if len(f.URIContact()) != 1 {
		t.Errorf("want 1 URI contact; have %d", len(f.URIContact()))
	}

	if len(f.EmailContact()) != 1 {
		t.Errorf("want 1 email contact; have %d", len(f.EmailContact()))
	}
}
