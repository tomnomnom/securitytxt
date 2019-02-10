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

	if len(f.Acknowledgments()) != 1 {
		t.Logf("Acknowledgments: %#v", f.Acknowledgments())
		t.Errorf("want 1 acknowledgments; have %d", len(f.Acknowledgments()))
	}

	if len(f.Contact()) != 3 {
		t.Logf("Contact: %#v", f.Contact())
		t.Errorf("want 3 contacts; have %d", len(f.Contact()))
	}

	if len(f.Encryption()) != 1 {
		t.Logf("Encryption: %#v", f.Encryption())
		t.Errorf("want 1 encryption; have %d", len(f.Encryption()))
	}

	if len(f.Hiring()) != 1 {
		t.Logf("Hiring: %#v", f.Hiring())
		t.Errorf("want 1 hiring; have %d", len(f.Hiring()))
	}

	if len(f.Policy()) != 1 {
		t.Logf("Policy: %#v", f.Policy())
		t.Errorf("want 1 policy; have %d", len(f.Policy()))
	}

	if len(f.PreferredLanguages()) != 1 {
		t.Logf("PreferredLanguages: %#v", f.PreferredLanguages())
		t.Errorf("want 1 preferred-languages; have %d", len(f.PreferredLanguages()))
	}

	if len(f.Comments()) != 3 {
		t.Logf("Comments: %#v", f.Comments())
		t.Errorf("want 3 comments; have %d", len(f.Comments()))
	}

	if len(f.Errors()) != 3 {
		t.Logf("Errors: %v", f.Errors())
		t.Errorf("want 3 errors; have %d", len(f.Errors()))
	}
}
