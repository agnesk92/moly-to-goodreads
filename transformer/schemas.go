package transformer

import (
	"encoding/xml"
	"fmt"
)

type MolyReads struct {
	Readings []Reading `xml:"reading"`
}

type Reading struct {
	AuthorAndTitle string `xml:"book"`

	Publisher          string `xml:"publisher"`
	YearOfPublication  string `xml:"year_of_publication"`
	PlaceOfPublication string `xml:"place_of_publication"`

	CreatedAt string `xml:"created_at"`
	StartTime string `xml:"start_time"`
	EndTime   string `xml:"end_time"`
	Earlier   string `xml:"earlier"`
	Stopped   string `xml:"stopped"`
	Pages     string `xml:"pages"`

	Note     string `xml:"note"`
	Comments string `xml:"comments"`
}

var goodReadsHeaders = []string{
	"Title",
	"Author",
	"ISBN",
	"My Rating",
	"Average Rating",
	"Publisher",
	"Binding",
	"Year Published",
	"Original Publication Year",
	"Date Read",
	"Date Added",
	"Shelves",
	"Bookshelves",
	"My Review",
}

func parseMolyData(xmldata []byte) (data MolyReads) {
	err := xml.Unmarshal(xmldata, &data)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(data.Readings); i++ {
		fmt.Println("Iteration : ", i)
		fmt.Println("From: " + data.Readings[i].AuthorAndTitle)
		fmt.Println()
	}

	return data
}
