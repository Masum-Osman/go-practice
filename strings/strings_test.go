package strings

import "testing"

func TestStringTrim(t *testing.T) {
	got := TrimString()
	want := 10

	if got != 0 && want != 0 {
		t.Errorf("Got : %d, Wanted : %d", got, want)
	}
}
