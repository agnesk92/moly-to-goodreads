package common

import (
	"encoding/csv"
	"io"
	"os"
)

func WriteCsv(csvData [][]string, csvPath string) {
	f, err := os.Create(csvPath)
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(f)
	err = writer.WriteAll(csvData)
	if err != nil {
		panic(err)
	}
}

func ReadXml(xmlPath string) []byte {
	xmlFile, err := os.Open(xmlPath)
	if err != nil {
		panic(err)
	}

	defer xmlFile.Close()

	data, err := io.ReadAll(xmlFile)
	if err != nil {
		panic(err)
	}

	return data
}
