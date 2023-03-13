package strings

import "testing"

func TestStringTrim(t *testing.T) {
	got := TrimString()
	want := "Hello, World"

	if got != want {
		t.Errorf("Got : %s, Wanted : %s", got, want)
	}
}
