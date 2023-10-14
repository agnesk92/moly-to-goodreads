package main

import (
	"flag"
	"moly-to-goodreads/transformer"
)

func main() {
	defaultInputPath := "./data/moly/readings.xml"
	defaultOutputPath := "./data/goodreads/readings.csv"

	inputPath := flag.String("sourcefile", defaultInputPath, "Source moly.hu xml file location")
	outputPath := flag.String("targetfile", defaultOutputPath, "Target goodreads csv file location")
	flag.Parse()

	transformer.MolyToGoodReads(
		*inputPath,
		*outputPath,
	)
}
