package transformer

import (
	"strings"

	"moly-to-goodreads/common"
)

func MolyToGoodReads(molyXmlPath, goodReadsTargetCsvPath string) {
	xmlData := common.ReadXml(molyXmlPath)
	molyData := parseMolyData(xmlData)

	goodReadsData := transform(molyData)

	common.WriteCsv(goodReadsData, goodReadsTargetCsvPath)
}

func transform(data MolyReads) [][]string {
	csvData := [][]string{}
	csvData = append(csvData, goodReadsHeaders)
	for i := 0; i < len(data.Readings); i++ {
		reading := data.Readings[i]

		author := ""
		title := reading.AuthorAndTitle
		if strings.Contains(reading.AuthorAndTitle, ":") {
			author = strings.Split(reading.AuthorAndTitle, ": ")[0]
			title = strings.Split(reading.AuthorAndTitle, ": ")[1]
		}

		createdAt := common.FormatDate(reading.CreatedAt)
		finishedAt := common.FormatDate(reading.EndTime)

		shelves := createShelves(finishedAt)

		row := createCsvRow(
			title,
			author,
			reading.Publisher,
			reading.YearOfPublication,
			finishedAt,
			createdAt,
			shelves,
		)
		csvData = append(csvData, row)
	}

	return csvData
}

func createShelves(finishedAt string) string {
	if finishedAt == "" {
		return "currently-reading"
	}
	return "read"
}

func createCsvRow(title, author, publisher, yearOfPublication, finishedAt, createdAt, shelves string) []string {
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
