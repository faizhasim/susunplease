package susun

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"regexp"
)

const (
	DocumentType = "documentType"
	TargetDir    = "targetDir"
	MatchRegex   = "matchRegex"
)

type Rule struct {
	DocumentType string
	TargetDir    string
	MatchRegex   *regexp.Regexp
}

type HeaderPos map[string]int

func headerPos(entries [][]string) (HeaderPos, error) {
	headerPos := make(HeaderPos, 3)
	for i, header := range entries[0] {
		switch header {
		case DocumentType:
			headerPos[DocumentType] = i
		case TargetDir:
			headerPos[TargetDir] = i
		case MatchRegex:
			headerPos[MatchRegex] = i
		default:
			return nil, errors.New(fmt.Sprintf("Unmatched csv header %s", header))
		}
	}
	return headerPos, nil
}

func entriesToRules(entries [][]string, pos HeaderPos) []Rule {
	var rules []Rule
	for _, entry := range entries[1:] {
		rule := Rule{
			DocumentType: entry[pos[DocumentType]],
			TargetDir:    entry[pos[TargetDir]],
			MatchRegex:   regexp.MustCompile("(?i)" + entry[pos[MatchRegex]]),
		}
		rules = append(rules, rule)
	}
	return rules
}

func ParseRulesFromCsv(path string) ([]Rule, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if fi.IsDir() {
		return nil, errors.New(fmt.Sprintf("%s is a directory.", path))
	}

	r := csv.NewReader(f)
	entries, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	pos, err := headerPos(entries)
	if err != nil {
		return nil, err
	}

	return entriesToRules(entries, pos), nil
}

func MatchRule(rules []Rule, content string) (Rule, bool) {
	for _, rule := range rules {
		if rule.MatchRegex.MatchString(content) {
			return rule, true
		}
	}
	return Rule{}, false
}
