package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "billy"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("billy")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("billy")) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")
	if message != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, message, err)
	}

}
