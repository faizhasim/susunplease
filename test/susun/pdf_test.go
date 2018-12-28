package susun_test

import (
	"github.com/faizhasim/susunplease/pkg/susun"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPdf(t *testing.T) {
	results, err := susun.ExtractPdfFromGlob("testdata/*")
	if err != nil {
		t.Error(err)
	}

	assert.Len(t, results, 2)

	for key, value := range results {
		t.Log(key, value)
		assert.Contains(t, value, "Title")
		assert.Contains(t, value, "Photo")
		switch key {
		case "testdata/profile1.pdf":
			assert.Contains(t, value, "Sample Content 1")
		case "testdata/profile2.pdf":
			assert.Contains(t, value, "Sample Content 2")
		default:
			t.Fail()
		}
	}
}
