package mailto

import (
	"testing"

	"github.com/team-ravl/civic-qa/services/mailto/internal/model"
)

func TestMailto(t *testing.T) {
	cases := []struct {
		config   model.Request
		expected string
	}{
		{model.Request{To: []string{"example@example.com"}}, `mailto:example@example.com`},
		{model.Request{To: []string{"harder+example@example.com"}}, `mailto:harder+example@example.com`},
		{model.Request{To: []string{"em1@mail.com", "em2@mail.com"}}, `mailto:em1@mail.com,em2@mail.com`},
		{model.Request{To: []string{"em1@mail.com"}, Cc: []string{"ccme@mail.com"}}, `mailto:em1@mail.com?cc=ccme@mail.com`},
		{model.Request{To: []string{"em1@mail.com"}, Bcc: []string{"bccme@mail.com"}}, `mailto:em1@mail.com?bcc=bccme@mail.com`},
		{model.Request{To: []string{"em1@mail.com"}, Cc: []string{"ccme@mail.com"}, Bcc: []string{"bccme@mail.com"}}, `mailto:em1@mail.com?cc=ccme@mail.com&bcc=bccme@mail.com`},
		{model.Request{To: []string{"em1@mail.com"}, Subject: "subj"}, `mailto:em1@mail.com?subject=subj`},
		{model.Request{To: []string{"em1@mail.com"}, Subject: "subj space"}, `mailto:em1@mail.com?subject=subj%20space`},
		{model.Request{To: []string{"em1@mail.com"}, Body: "body"}, `mailto:em1@mail.com?body=body`},
		{model.Request{To: []string{"em1@mail.com"}, Body: "body space"}, `mailto:em1@mail.com?body=body%20space`},
		{model.Request{To: []string{"em1@mail.com"}, Subject: "subject $ (wow)!", Body: "break\nbreak"}, `mailto:em1@mail.com?subject=subject%20$%20%28wow%29%21&body=break%0Abreak`},
	}

	for i, testCase := range cases {
		output, err := Generate(testCase.config)
		if err != nil {
			t.Fatalf("(%d) failed to generate: %v", i, err)
		}

		if string(output) != testCase.expected {
			t.Fatalf("(%d) Match failed: got %s, expected %s", i, output, testCase.expected)
		}
	}
}
