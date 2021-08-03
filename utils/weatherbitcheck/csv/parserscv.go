package csv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func CSVToMap(filepath string) (*map[string]string, error) {

	csvFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()


	r := csv.NewReader(bufio.NewReader(csvFile))
	r.Comma = ';'
	rows := map[string]string{}


	index := 0
	for {
		index++
		line, error := r.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			fmt.Println("err", err)
			return nil, err
		}
		if index == 1 {
			continue
		}
		innerID := line[0]
		weatherBitID := line[1]
		rows[innerID] = weatherBitID
	}

	return &rows, nil
}