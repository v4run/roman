package roman

import (
	"testing"
)

var arabicSuccessTests = []struct {
	input    string
	expected int
}{
	{"I", 1},
	{"XIX", 19},
	{"MCMIII", 1903},
	{"XLIII", 43},
	{"CIX", 109},
	{"CCCXCVII", 397},
	{"MMMCMXCIX", 3999},
}

var arabicErrorTests = []struct {
	input string
}{
	{"XXXX"},
	{"CCCD"},
	{"IXI"},
	{""},
	{"INVALID"},
	{"DM"},
	{"DD"},
	{"IIII"},
	{"MMMM"},
	{"CCCC"},
	{"IM"},
	{"XDI"},
}

var romanTests = []struct {
	input    int
	expected string
}{
	{1, "I"},
	{19, "XIX"},
	{1903, "MCMIII"},
	{3999, "MMMCMXCIX"},
	{43, "XLIII"},
	{397, "CCCXCVII"},
}

func TestNew(t *testing.T) {
	t.Log("\nTesting New() | Success cases\n")
	for _, tc := range arabicSuccessTests {
		r, _ := New(tc.input)
		if r.Arabic() != tc.expected {
			t.Errorf("Failure | New(%s): expected = `%d`, actual = `%d`\n", tc.input, tc.expected, r.Arabic())
		} else {
			t.Logf("Success | New(%s): expected = `%d`, actual = `%d`\n", tc.input, tc.expected, r.Arabic())
		}
	}
	t.Log("\nTesting New() | Error cases\n")
	for _, tc := range arabicErrorTests {
		if _, err := New(tc.input); err == nil {
			t.Errorf("Failure | New(%s): expected = error, actual = `%v`", tc.input, err)
		} else {
			t.Logf("Success | New(%s): expected = `%s`, actual = `%s`", tc.input, err, err)
		}

	}
}

func TestFromArabic(t *testing.T) {
	t.Log("\nTesting FromArabic()\n")
	for _, tc := range romanTests {
		actual := FromArabic(tc.input)
		if actual.Roman() != tc.expected {
			t.Errorf("Failure | FromArabic(%d): expected = `%s`, actual = `%s`", tc.input, tc.expected, actual.Roman())
		} else {
			t.Logf("Success | FromArabic(%d): expected = `%s`, actual = `%s`", tc.input, tc.expected, actual.Roman())
		}
	}
}
