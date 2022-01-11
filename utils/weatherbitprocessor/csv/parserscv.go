package csv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

type TempData struct {
	Date string
	Temp float64
	Timezone string
}

func CSVToTempsData(filepath string) (*[]TempData, error) {

	csvFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()


	r := csv.NewReader(bufio.NewReader(csvFile))
	r.Comma = ';'
	rows := make([]TempData, 0)

	index := 0
	for {
		index++
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("err", err)
			return nil, err
		}
		if index == 1 {
			fmt.Println("line", line)
			continue
		}

		t := strings.Replace(line[4], ",", ".", 1)
		temp, err := strconv.ParseFloat(t, 10)
		if err != nil {
			fmt.Println("temp parse err", err)
			continue
		}

		data := line[0]
		timezone := line[1]
		rows = append(rows, TempData{
			data,
			temp,
			timezone,
		})
	}

	return &rows, nil
}