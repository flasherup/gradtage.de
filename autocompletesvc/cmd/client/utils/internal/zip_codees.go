package internal

import (
	"encoding/csv"
	"github.com/flasherup/gradtage.de/autocompletesvc"
	"io"
	"os"
)

func LoadZipcodes(path string) ([]autocompletesvc.ZipCode, error) {
	csvFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	r := csv.NewReader(csvFile)
	r.Comma = ','
	zipCodes := make([]autocompletesvc.ZipCode, 0)

	index := 0
	for {
		index++
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if index == 1 {
			continue
		}

		zipCodes = append(zipCodes, autocompletesvc.ZipCode{
			ZipCode: line[0],
			Station: line[1],
		})
	}

	return zipCodes, nil
}

func FormatZipCode(codes []autocompletesvc.ZipCode) {
	for i, v := range codes {
		prefix := v.Station[:3]
		codes[i].ZipCode = prefix + v.ZipCode
	}
}
