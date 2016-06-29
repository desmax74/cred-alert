package git

import (
	"cred-alert/patterns"

	"github.com/pivotal-golang/lager"
)

type Scanner interface {
	Scan(lager.Logger) bool
	Line() *Line
}

func Sniff(logger lager.Logger, scanner Scanner) []Line {
	logger = logger.Session("sniff")

	matcher := patterns.DefaultMatcher()

	matchingLines := []Line{}

	for scanner.Scan(logger) {
		line := *scanner.Line()
		found := matcher.Match(line.Content)

		if found {
			matchingLines = append(matchingLines, line)
		}
	}

	return matchingLines
}
