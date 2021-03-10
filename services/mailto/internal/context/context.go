package context

import "github.com/vivian-hua/civic-qa/services/common/config"

// Context containers handler context information (none currently)
type Context struct {
	// placeholder for handler context items
}

// BuildContext builds a handler Context based on a config.Provider
// (does nothing currently)
func BuildContext(cfg config.Provider) (*Context, error) {
	return &Context{}, nil
}
