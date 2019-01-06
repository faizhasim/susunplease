package model

import "regexp"

type Rule struct {
	DocumentType string
	TargetDir    string
	MatchRegex   *regexp.Regexp
}
