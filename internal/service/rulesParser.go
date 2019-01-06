package service

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/faizhasim/susunplease/internal/model"
	"github.com/mitchellh/go-homedir"
	"os"
	"regexp"
)

const (
	DocumentType = "documentType"
	TargetDir    = "targetDir"
	MatchRegex   = "matchRegex"
)

type RulesParser interface {
	ParseRulesFromCsv(path string) ([]model.Rule, error)
	MatchRule(rules []model.Rule, content string) (model.Rule, bool)
	GetCsvPath() (string, error)
}

type rulesParser struct{}

func headerPos(entries [][]string) (map[string]int, error) {
	headerPos := make(map[string]int, 3)
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

func entriesToRules(entries [][]string, pos map[string]int) []model.Rule {
	var rules []model.Rule
	for _, entry := range entries[1:] {
		rule := model.Rule{
			DocumentType: entry[pos[DocumentType]],
			TargetDir:    entry[pos[TargetDir]],
			MatchRegex:   regexp.MustCompile("(?i)" + entry[pos[MatchRegex]]),
		}
		rules = append(rules, rule)
	}
	return rules
}

func (rulesParser *rulesParser) ParseRulesFromCsv(path string) ([]model.Rule, error) {
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

func (rulesParser *rulesParser) MatchRule(rules []model.Rule, content string) (model.Rule, bool) {
	for _, rule := range rules {
		if rule.MatchRegex.MatchString(content) {
			return rule, true
		}
	}
	return model.Rule{}, false
}

func (rulesParser *rulesParser) GetCsvPath() (string, error) {
	h, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return h + "/.susun/rules.csv", nil
}

func NewRulesParser() RulesParser {
	return &rulesParser{}
}
