package transformer

import (
	"strings"

	"moly-to-goodreads/common"
)

func MolyToGoodReads(molyXMLPath, goodReadsTargetCSVPath string) {
	xmlData := common.ReadXML(molyXMLPath)
	molyData := parseMolyData(xmlData)

	goodReadsData := transform(molyData)

	common.WriteCSV(goodReadsData, goodReadsTargetCSVPath)
}

func transform(data MolyReads) [][]string {
	csvData := [][]string{}
	csvData = append(csvData, goodReadsHeaders)
	for i := 0; i < len(data.Readings); i++ {
		csvData = append(
			csvData,
			transformReading(&data.Readings[i]),
		)
	}

	return csvData
}

func transformReading(reading *Reading) []string {
	author := ""
	title := reading.AuthorAndTitle
	if strings.Contains(reading.AuthorAndTitle, ":") {
		author = strings.Split(reading.AuthorAndTitle, ": ")[0]
		title = strings.Split(reading.AuthorAndTitle, ": ")[1]
	}

	createdAt := common.FormatDate(reading.CreatedAt)
	finishedAt := common.FormatDate(reading.EndTime)

	shelves := createShelves(finishedAt)

	return createCsvRow(
		title,
		author,
		reading.Publisher,
		reading.YearOfPublication,
		finishedAt,
		createdAt,
		shelves,
	)
}

func createShelves(finishedAt string) string {
	if finishedAt == "" {
		return "currently-reading"
	}
	return "read"
}

func createCsvRow(title,
	author,
	publisher,
	yearOfPublication,
	finishedAt,
	createdAt,
	shelves string,
) []string {
	return []string{
		title,
		author,
		"", "", "",
		publisher,
		"",
		yearOfPublication,
		"",
		finishedAt,
		createdAt,
		shelves,
		"", "",
	}
}
