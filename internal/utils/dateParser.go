package utils

import (
	"errors"
	"regexp"
	"sort"
	"strings"
	"time"
)

var layouts []string

func init() {
	layouts = []string{
		"January 2, 2006",
		"2 January, 2006",
		"2 January 2006",
		"Jan 2, 2006",
		"2 Jan, 2006",
		"2 Jan 2006",
		"02/01/06",
		"02-01-06",
		"02.01.06",
		"02 01 06",
		"02/01/2006",
		"02-01-2006",
		"02.01.2006",
		"02 01 2006",
		"Jan/02/06",
		"Jan-02-06",
		"Jan 02 06",
		"02/Jan/06",
		"02-Jan-06",
		"02 Jan 06",
		"January/02/06",
		"January-02-06",
		"January 02 06",
		"02/January/06",
		"02-January-06",
		"02 January 06",
		"2006-01-02",
		"2006.01.02",
		"2006 01 02",
	}
}

func ParseTimes(txt string) (time.Time, error) {
	times := make([]time.Time, 0)
	for _, layout := range layouts {

		pattern := strings.Replace(layout, "January", "\\w+", -1)
		pattern = strings.Replace(pattern, "Jan", "\\w+", -1)
		pattern = strings.Replace(pattern, "2006", "[0-9]{4}", -1)
		pattern = strings.Replace(pattern, "06", "[0-9]{2}", -1)
		pattern = strings.Replace(pattern, "02", "[0-9]{2}", -1)
		pattern = strings.Replace(pattern, "01", "[0-9]{2}", -1)

		regex := regexp.MustCompile(pattern)

		matches := regex.FindAllString(txt, -1)
		for _, match := range matches {
			ts, err := time.Parse(layout, match)
			if err != nil {
				// log.Println(fmt.Sprintf("Unable to parse using layout %s: %s", layout, match))
			}
			timeLowerLimit, err := time.Parse("2006-01-02", "2000-01-01")
			if err != nil {
				panic(err)
			}

			if ts.After(time.Now()) {
				// log.Println(fmt.Sprintf("Excluding date after now: %s", ts.Format("2006-01-02")))
			} else if ts.Before(timeLowerLimit) {
				// log.Println(fmt.Sprintf("Excluding date after %s: %s", timeLowerLimit.Format("2006-01-02"), ts.Format("2006-01-02")))
			} else {
				times = append(times, ts)
			}
		}
	}

	if len(times) == 0 {
		return time.Time{}, errors.New("no time can be parsed")
	}

	sort.Slice(times, func(i, j int) bool {
		return times[i].After(times[j])
	})

	return times[0], nil

}
