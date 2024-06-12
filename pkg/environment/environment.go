package environment

import "strings"

type Environment string

const (
	Production  Environment = "production"
	Test        Environment = "test"
	Development Environment = "development"
	Local       Environment = "local"
)

func New(name string) Environment {
	switch strings.ToLower(name) {
	case "production", "prod", "p":
		return Production
	case "test", "t":
		return Test
	case "development", "dev", "d":
		return Development
	case "local", "l":
		return Local
	default:
		return Local
	}
}
