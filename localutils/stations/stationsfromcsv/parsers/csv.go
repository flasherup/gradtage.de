package parsers

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/flasherup/gradtage.de/autocompletesvc"
	"github.com/flasherup/gradtage.de/stationssvc"
	"io"
	"os"
	"strings"
)

func CSVToStationsList(filepath string) ([]stationssvc.Station, error) {
	fmt.Println("filepath", filepath)
	csvFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()


	r := csv.NewReader(bufio.NewReader(csvFile))
	r.Comma = ','
	stations := make([]stationssvc.Station, 0)


	index := 0
	for {
		index++
		line, error := r.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			return nil, err
		}
		if index == 1 {
			continue
		}
		innerID := line[0]
		weatherBitID := line[1]
		cityName := line[9]
		cityName = strings.Replace(cityName, "'", "''", -1)
		stations = append(stations, stationssvc.Station{
			ID:innerID,
			Name:cityName,
			SourceID: weatherBitID,
		})
	}

	return stations, nil
}

func CSVToAutocompleteList(filepath string) ([]autocompletesvc.Source, error) {
	fmt.Println("filepath", filepath)
	csvFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()

	r := csv.NewReader(bufio.NewReader(csvFile))
	r.Comma = ';'
	stations := make([]autocompletesvc.Source, 0)


	index := 0
	for {
		index++
		line, error := r.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			return nil, err
		}
		if index == 1 {
			continue
		}
		innerID := line[0]
		cityName := line[9]
		icao := validateId(line[12])
		wmo := validateId(line[13])
		cwop := validateId(line[14])
		cityName = strings.Replace(cityName, "'", "''", -1)
		fmt.Println("cityName", cityName)
		stations = append(stations, autocompletesvc.Source {
			ID:innerID,
			Name:cityName,
			Icao: icao,
			Dwd: "",
			Wmo: wmo,
			Cwop: cwop,
		})
	}

	return stations, nil
}

func validateId(src string) string {
	if src == "0" {
		return ""
	}

	return src
}
