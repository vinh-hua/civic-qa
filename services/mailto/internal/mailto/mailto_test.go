package mailto

import (
	"testing"

	"github.com/vivian-hua/civic-qa/services/mailto/internal/model"
)

func TestMailto(t *testing.T) {
	cases := []struct {
		config   model.Request
		expected string
	}{
		{model.Request{To: []string{"example@example.com"}, InnerText: "click"}, `<a href="mailto:example@example.com">click</a>`},
		{model.Request{To: []string{"harder+example@example.com"}, InnerText: "click"}, `<a href="mailto:harder+example@example.com">click</a>`},
		{model.Request{To: []string{"em1@mail.com", "em2@mail.com"}, InnerText: "a"}, `<a href="mailto:em1@mail.com,em2@mail.com">a</a>`},
		{model.Request{To: []string{"em1@mail.com"}, Cc: []string{"ccme@mail.com"}, InnerText: "a"}, `<a href="mailto:em1@mail.com?cc=ccme@mail.com">a</a>`},
		{model.Request{To: []string{"em1@mail.com"}, Bcc: []string{"bccme@mail.com"}, InnerText: "a"}, `<a href="mailto:em1@mail.com?bcc=bccme@mail.com">a</a>`},
		{model.Request{To: []string{"em1@mail.com"}, Cc: []string{"ccme@mail.com"}, Bcc: []string{"bccme@mail.com"}, InnerText: "a"}, `<a href="mailto:em1@mail.com?cc=ccme@mail.com&bcc=bccme@mail.com">a</a>`},
		{model.Request{To: []string{"em1@mail.com"}, Subject: "subj", InnerText: "a"}, `<a href="mailto:em1@mail.com?subject=subj">a</a>`},
		{model.Request{To: []string{"em1@mail.com"}, Subject: "subj space", InnerText: "a"}, `<a href="mailto:em1@mail.com?subject=subj%20space">a</a>`},
		{model.Request{To: []string{"em1@mail.com"}, Body: "body", InnerText: "a"}, `<a href="mailto:em1@mail.com?body=body">a</a>`},
		{model.Request{To: []string{"em1@mail.com"}, Body: "body space", InnerText: "a"}, `<a href="mailto:em1@mail.com?body=body%20space">a</a>`},
		{model.Request{To: []string{"em1@mail.com"}, Subject: "subject $ (wow)!", Body: "break\nbreak"}, `<a href="mailto:em1@mail.com?subject=subject%20$%20%28wow%29%21&body=break%0Abreak"></a>`},
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
