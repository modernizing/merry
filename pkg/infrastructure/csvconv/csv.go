package csvconv

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func ParseCSV(path string) [][]string {
	var results [][]string
	csvfile, err := os.Open(path)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)
	for {
		var result []string
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		for _, rec := range record {
			result = append(result, rec)
		}
		results = append(results, result)
	}
	return results
}
