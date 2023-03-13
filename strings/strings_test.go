package strings

import (
	"testing"
)

type testCases struct {
	arg  string
	want bool
}

func TestString(t *testing.T) {

	t.Run("Is Only Letters tests: ", func(t *testing.T) {
		cases := []testCases{
			{"Golang", true},
			{"Golang542", false},
			{"Golang542", false},
		}

		for _, tc := range cases {
			got := IsOnlyLetter(tc.arg)
			if tc.want != got {
				t.Errorf("Expected '%t', but got '%t'", tc.want, got)
			}
		}
	})

	t.Run("String Trimming Test:", func(t *testing.T) {
		got := TrimString()
		want := "Hello, World"

		if got != want {
			t.Errorf("Got : %s, Wanted : %s", got, want)
		}
	})

	t.Run("Is All Upper case and only Letter test :", func(t *testing.T) {
		cases := []testCases{
			{"UPPERCASE", true},
			{"LowerCase", false},
		}

		for _, tc := range cases {
			got := AllUpperCase(tc.arg)

			if tc.want != got {
				t.Errorf("Expected '%t', but got '%t'", tc.want, got)
			}
		}
	})

}
