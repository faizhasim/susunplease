package susun_test

import (
	"github.com/faizhasim/susunplease/pkg/susun"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestReadingRules(t *testing.T) {
	records, err := susun.ParseRulesFromCsv("testdata/rules.csv")
	if err != nil {
		t.Error(err)
	}
	assert.Len(t, records, 1)
	assert.Equal(t, "sugar-high-inc", records[0].DocumentType)
	assert.Equal(t, "receipt/food", records[0].TargetDir)
	assert.Equal(t, regexp.MustCompile("(?i)sugar.*high"), records[0].MatchRegex)
}
