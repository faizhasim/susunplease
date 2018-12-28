package susun_test

import (
	"github.com/faizhasim/susunplease/pkg/susun"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func testFromSample(t *testing.T, sample string, expectedIsoDate string) {
	expectNoResults := expectedIsoDate == ""
	expected, err := time.Parse("2006-01-02", expectedIsoDate)
	if err != nil && !expectNoResults {
		panic(err)
	}

	tm, err := susun.ParseTimes(sample)
	if err != nil {
		if expectNoResults {
			assert.Error(t, err)
		} else {
			panic(err)
		}
	}

	assert.Equal(t, expected, tm)
}

func TestDateParsing(t *testing.T) {
	testFromSample(t, "today is asd", "")
	testFromSample(t, "today is 2018-12-23, not 2018-12-21", "2018-12-23")
	testFromSample(t, "today is 23-12-2018, not 21-12-2017", "2018-12-23")
	testFromSample(t, "today is 23.12.2018, not 21.12.2017", "2018-12-23")
}
