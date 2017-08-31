package securitytxt

import "testing"

func TestFilters(t *testing.T) {
	f := &File{}
	f.addContact("https://example.com")
	f.addContact("+44 5555 555 555")
	f.addContact("mail@example.com")
	f.addDisclosure("full")

	if len(f.PhoneContact()) != 1 {
		t.Errorf("want 1 phone contact; have %d", len(f.PhoneContact()))
	}

	if len(f.URIContact()) != 1 {
		t.Errorf("want 1 URI contact; have %d", len(f.URIContact()))
	}

	if len(f.EmailContact()) != 1 {
		t.Errorf("want 1 email contact; have %d", len(f.EmailContact()))
	}

	if f.HasFullDisclosure() != true {
		t.Errorf("want true for HasFullDisclosure(); have false")
	}
}
