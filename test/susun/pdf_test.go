package susun_test

import (
	"github.com/faizhasim/susunplease/internal/module"
	"github.com/faizhasim/susunplease/internal/sideEffect"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPdf1(t *testing.T) {
	m := &module.SusunPdf{PdfService: sideEffect.NewPdfService()}

	content, err := m.ParsePdfContent("testdata/profile1.pdf")
	if err != nil {
		t.Error(err)
	}
	assert.Contains(t, content, "Title")
	assert.Contains(t, content, "Photo")
	assert.Contains(t, content, "Sample Content 1")
}
