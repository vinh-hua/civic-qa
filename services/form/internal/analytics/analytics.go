package analytics

import "github.com/team-ravl/civic-qa/services/form/internal/model"

type Client interface {
	GetKeyPhrases(resp *model.FormResponse) ([]string, error)
}

type AnalyticsResponse = []string
