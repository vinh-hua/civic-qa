package model

import "testing"

func TestVerifyEmail(t *testing.T) {
	cases := []struct {
		email      string
		shouldPass bool
	}{
		{"example@example.com", true},
		{"example+plus@example.com", true},
		{"_____@example.com", true},
		{"email@subdomain.example.com", true},
		{"bad email@example.com", false},
		{"bad@-.com", false},
		{"@example.com", false},
		{"plain", false},
		{"email.example.com", false},
	}

	for i, testCase := range cases {
		if verifyEmail(testCase.email) != testCase.shouldPass {
			t.Fatalf("(%d) %s expected %t", i, testCase.email, testCase.shouldPass)
		}
	}
}
