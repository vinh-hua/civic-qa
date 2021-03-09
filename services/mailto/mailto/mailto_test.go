package mailto

import "testing"

func TestGenerator(t *testing.T) {
	cfg := ReplyConfig(
		[]string{"easy0117@gmail.com"},
		nil,
		nil,
		"subject1",
		"body2",
		"click",
	)
	o, _ := Generate(cfg)
	t.Log(o)
}

func TestMailto(t *testing.T) {
	cases := []struct {
		config   Config
		expected string
	}{
		{Config{To: []string{"example@example.com"}, InnerText: "click"}, `<a href="mailto:example@example.com">click</a>`},
		{Config{To: []string{"harder+example@example.com"}, InnerText: "click"}, `<a href="mailto:harder+example@example.com">click</a>`},
		{Config{To: []string{"em1@mail.com", "em2@mail.com"}, InnerText: "a"}, `<a href="mailto:em1@mail.com,em2@mail.com">a</a>`},
		{Config{To: []string{"em1@mail.com"}, Cc: []string{"ccme@mail.com"}, InnerText: "a"}, `<a href="mailto:em1@mail.com?cc=ccme@mail.com">a</a>`},
		{Config{To: []string{"em1@mail.com"}, Bcc: []string{"bccme@mail.com"}, InnerText: "a"}, `<a href="mailto:em1@mail.com?bcc=bccme@mail.com">a</a>`},
		{Config{To: []string{"em1@mail.com"}, Cc: []string{"ccme@mail.com"}, Bcc: []string{"bccme@mail.com"}, InnerText: "a"}, `<a href="mailto:em1@mail.com?cc=ccme@mail.com&bcc=bccme@mail.com">a</a>`},
		{Config{To: []string{"em1@mail.com"}, Subject: "subj", InnerText: "a"}, `<a href="mailto:em1@mail.com?subject=subj">a</a>`},
		{Config{To: []string{"em1@mail.com"}, Subject: "subj space", InnerText: "a"}, `<a href="mailto:em1@mail.com?subject=subj%20space">a</a>`},
		{Config{To: []string{"em1@mail.com"}, Body: "body", InnerText: "a"}, `<a href="mailto:em1@mail.com?body=body">a</a>`},
		{Config{To: []string{"em1@mail.com"}, Body: "body space", InnerText: "a"}, `<a href="mailto:em1@mail.com?body=body%20space">a</a>`},
		{Config{To: []string{"em1@mail.com"}, Subject: "subject $ (wow)!", Body: "break\nbreak"}, `<a href="mailto:em1@mail.com?subject=subject%20$%20%28wow%29%21&body=break%0Abreak"></a>`},
	}

	for i, testCase := range cases {
		output, err := Generate(testCase.config)
		if err != nil {
			t.Fatalf("(%d) failed to generate: %v", i, err)
		}

		if output != testCase.expected {
			t.Fatalf("(%d) Match failed: got %s, expected %s", i, output, testCase.expected)
		}
	}
}

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