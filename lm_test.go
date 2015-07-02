package lm

import "testing"

var (
	ErrMissingURL = &Message{Level: LevelError, Code: 100, Name: "missing_url", Value: "Missing URL"}
)

func TestSimpleError(t *testing.T) {
	if expect, got := `ERROR |  | 100 | missing_url |  | ["Missing URL"]`, ErrMissingURL.String(); expect != got {
		t.Fatalf("Unexpected output:\nExpect:\t%s\nGot:\t%s", expect, got)
	}
}

func TestNestedError(t *testing.T) {
	if expect, got := `ERROR | lm_test.go:16 | 100 | missing_url |  | ["Missing URL"]`, NewError(ErrMissingURL).String(); expect != got {
		t.Fatalf("Unexpected output:\nExpect:\t%s\nGot:\t%s", expect, got)
	}
}

func TestNestedMessage(t *testing.T) {
	dump := ErrMissingURL.NewMessagef("hello").String()
	if expect, got := `INFO | lm_test.go:22 | 100 | missing_url |  | ["hello","Missing URL"]`, string(dump); expect != got {
		t.Fatalf("Unexpected output:\nExpect:\t%s\nGot:\t%s", expect, got)
	}
}
