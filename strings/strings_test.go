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

	t.Run("If word contains: ", func(t *testing.T) {

		type tests struct {
			word   string
			letter string
			result bool
		}

		cases := []tests{
			{"abcd", "b", true},
			{"masum", "m", true},
			{"mishqat", "m", true},
			{"vscode", "w", false},
			{"vscode", "w", true},
			{"vscode", "w", true},
			{"vscode", "se", true},
		}

		for _, tc := range cases {
			got := IfContains(tc.word, tc.letter)

			if tc.result != got {
				t.Errorf("Expected '%t', but got '%t'", tc.result, got)
			}
		}
	})

}

func TestDefangIPaddr(t *testing.T) {
	t.Run("1108. Defanging an IP Address: ", func(t *testing.T) {
		type defangTestCases struct {
			param  string
			result string
		}

		cases := []defangTestCases{
			{"1.1.1.1", "1[.]1[.]1[.]1"},
			{"255.100.50.0", "255[.]100[.]50[.]0"},
		}

		for _, tc := range cases {
			got := DefangIPaddr(tc.param)

			t.Logf("Got : %s, Wanted : %s", got, tc.result)

			if tc.result != got {
				t.Errorf("Got : %s, Wanted : %s", got, tc.result)
			}
		}
	})
}

func TestFinalValue(t *testing.T) {
	t.Run("2011. Final Value of Variable After Performing Operations", func(t *testing.T) {

		type testCases struct {
			ops    []string
			result int32
		}

		cases := []testCases{
			{[]string{"++X", "++X", "X++"}, 3},
			{[]string{"--X", "X++", "X++"}, 1},
			{[]string{"X++", "++X", "--X", "X--"}, 0},
		}

		for _, tc := range cases {
			got := FinalValueAfterOperations(tc.ops)
			t.Log(got)

			if int(tc.result) != got {
				t.Errorf("Got : %d, Wanted : %d", got, int(tc.result))
			}
		}

	})
}
