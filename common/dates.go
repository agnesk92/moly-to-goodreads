package common

import (
	"errors"
	"strings"
)

var mapMonth = map[string]string{
	"január":     "01",
	"február":    "02",
	"március":    "03",
	"április":    "04",
	"május":      "05",
	"június":     "06",
	"július":     "07",
	"augusztus":  "08",
	"szeptember": "09",
	"október":    "10",
	"november":   "11",
	"december":   "12",
}

func FormatDate(date string) string {
	splitted := strings.Split(date, " ")
	if len(splitted) < 3 {
		return ""
	}

	month := cleanString(splitted[1])
	year := cleanString(splitted[0])

	day := splitted[2]
	if day == "" {
		day = "0" + splitted[3]
	}
	day = cleanString(day)

	mappedMonth, ok := mapMonth[month]
	if !ok {
		panic(errors.New("Month not found: " + month))
	}

	return year + "-" + mappedMonth + "-" + day
}

func cleanString(dirty string) string {
	cleanse := strings.NewReplacer(
		".", "",
		",", "",
	)
	return cleanse.Replace(dirty)
}
