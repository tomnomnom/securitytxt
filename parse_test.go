package securitytxt

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	r, err := os.Open("test/basic.txt")
	if err != nil {
		t.Fatalf(err.Error())
	}

	f, err := FromReader(r)

	if err == nil {
		t.Logf("want non-nil error; have nil")
	}

	if len(f.Comments()) != 3 {
		t.Logf("Comments: %#v", f.Comments())
		t.Errorf("want 3 comments; have %d", len(f.Comments()))
	}

	if len(f.Errors()) != 4 {
		t.Logf("Errors: %v", f.Errors())
		t.Errorf("want 4 errors; have %d", len(f.Errors()))
	}

	if len(f.Contact()) != 3 {
		t.Logf("Contact: %#v", f.Contact())
		t.Errorf("want 3 contacts; have %d", len(f.Contact()))
	}

	if len(f.Disclosure()) != 3 {
		t.Logf("Disclosure: %#v", f.Disclosure())
		t.Errorf("want 3 disclosures; have %d", len(f.Disclosure()))
	}

	if len(f.Encryption()) != 1 {
		t.Logf("Encryption: %#v", f.Encryption())
		t.Errorf("want 1 encryption; have %d", len(f.Encryption()))
	}

	if len(f.Acknowledgement()) != 1 {
		t.Logf("Acknowledgement: %#v", f.Acknowledgement())
		t.Errorf("want 1 acknowledgement; have %d", len(f.Acknowledgement()))
	}
}
