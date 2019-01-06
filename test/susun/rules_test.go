package susun_test

import (
	"github.com/faizhasim/susunplease/internal/module"
	"github.com/faizhasim/susunplease/internal/service"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestReadingRules(t *testing.T) {
	m := &module.SusunPdf{RulesParser: service.NewRulesParser()}
	records, err := m.ParseRulesFromCsv("testdata/rules.csv")
	if err != nil {
		t.Error(err)
	}
	assert.Len(t, records, 1)
	assert.Equal(t, "sugar-high-inc", records[0].DocumentType)
	assert.Equal(t, "receipt/food", records[0].TargetDir)
	assert.Equal(t, regexp.MustCompile("(?i)sugar.*high"), records[0].MatchRegex)
}
