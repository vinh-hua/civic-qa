package parse

import "testing"

func TestParseUintOrDefault(t *testing.T) {
	cases := []struct {
		str string
		exp uint
	}{
		{
			str: "1",
			exp: 1,
		},
		{
			str: "0",
			exp: 0,
		},
		{
			str: "10",
			exp: 10,
		},
		{
			// defaults
			str: "a",
			exp: 0,
		},
		{
			// defaults
			str: "notnum",
			exp: 0,
		},
		{
			// defaults
			str: "",
			exp: 0,
		},
	}

	for i, testCase := range cases {
		out := ParseUintOrDefault(testCase.str)
		if out != testCase.exp {
			t.Fatalf("(%d) Expected: %v, got: %v", i, testCase.exp, out)
		}
	}
}

func TestParseBoolOrDefault(t *testing.T) {
	cases := []struct {
		str string
		exp bool
	}{
		{
			str: "true",
			exp: true,
		},
		{
			str: "TRUE",
			exp: true,
		},
		{
			str: "True",
			exp: true,
		},
		{
			str: "t",
			exp: true,
		},
		{
			str: "T",
			exp: true,
		},
		{
			str: "1",
			exp: true,
		},
		{
			str: "false",
			exp: false,
		},
		{
			str: "FALSE",
			exp: false,
		},
		{
			str: "False",
			exp: false,
		},
		{
			str: "f",
			exp: false,
		},
		{
			str: "F",
			exp: false,
		},
		{
			str: "0",
			exp: false,
		},
		{
			// defaults
			str: "2",
			exp: false,
		},
		{
			// defaults
			str: "bad",
			exp: false,
		},
		{
			// defaults
			str: "",
			exp: false,
		},
		{
			// defaults
			str: "tf",
			exp: false,
		},
		{
			// defaults
			str: "yes",
			exp: false,
		},
	}

	for i, testCase := range cases {
		out := ParseBoolOrDefault(testCase.str)
		if out != testCase.exp {
			t.Fatalf("(%d) Expected: %v, got: %v", i, testCase.exp, out)
		}
	}
}
