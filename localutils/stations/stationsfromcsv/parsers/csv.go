package parsers

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/flasherup/gradtage.de/stationssvc"
	"io"
	"os"
	"strconv"
	"strings"
)

type Station struct {
	ID string
	SourceID string
	Latitude float64
	Longitude float64
	Source string
	Reports string
	ISO2Country string
	ISO3Country string
	Prio string
	CityNameEnglish string
	CityNameNative string
	CountryNameEnglish string
	CountryNameNative string
	ICAO string
	WMO string
	CWOP string
	Maslib string
	National_ID string
	IATA string
	USAF_WBAN string
	GHCN string
	NWSLI string
	Elevation float64
}

func ParseStationsCSV(filepath string) ([]Station, error) {
	fmt.Println(filepath)
	csvFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	r := csv.NewReader(bufio.NewReader(csvFile))
	r.Comma = ','
	stations := make([]Station, 0)


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
		stations = append(stations, Station{
			ID:line[0],
			SourceID:line[1],
			Latitude:prepareFloat64(line[2]),
			Longitude:prepareFloat64(line[3]),
			Source:line[4],
			Reports:line[5],
			ISO2Country:line[6],
			ISO3Country:line[7],
			Prio:line[8],
			CityNameEnglish:prepareName(line[9]),
			CityNameNative:prepareName(line[10]),
			CountryNameEnglish:prepareName(line[11]),
			CountryNameNative:prepareName(line[12]),
			ICAO:line[13],
			WMO:line[14],
			CWOP:line[15],
			Maslib:line[16],
			National_ID:line[17],
			IATA:line[18],
			USAF_WBAN:line[19],
			GHCN:line[20],
			NWSLI:line[21],
			Elevation:prepareFloat64(line[22]),
		})
	}
	return stations, nil
}

func prepareName(name string) string  {
	n := fromWindows1252(name)
	if len(name) < 2 {
		return name
	}
	return strings.Replace(n, "'", "''", -1)
}

func fromWindows1252(str string) string {
	var arr = []byte(str)
	var buf bytes.Buffer
	var r rune

	for _, b := range(arr) {
		switch b {
		case 0x80:
			r = 0x20AC
		case 0x82:
			r = 0x201A
		case 0x83:
			r = 0x0192
		case 0x84:
			r = 0x201E
		case 0x85:
			r = 0x2026
		case 0x86:
			r = 0x2020
		case 0x87:
			r = 0x2021
		case 0x88:
			r = 0x02C6
		case 0x89:
			r = 0x2030
		case 0x8A:
			r = 0x0160
		case 0x8B:
			r = 0x2039
		case 0x8C:
			r = 0x0152
		case 0x8E:
			r = 0x017D
		case 0x91:
			r = 0x2018
		case 0x92:
			r = 0x2019
		case 0x93:
			r = 0x201C
		case 0x94:
			r = 0x201D
		case 0x95:
			r = 0x2022
		case 0x96:
			r = 0x2013
		case 0x97:
			r = 0x2014
		case 0x98:
			r = 0x02DC
		case 0x99:
			r = 0x2122
		case 0x9A:
			r = 0x0161
		case 0x9B:
			r = 0x203A
		case 0x9C:
			r = 0x0153
		case 0x9E:
			r = 0x017E
		case 0x9F:
			r = 0x0178
		default:
			r = rune(b)
		}

		buf.WriteRune(r)
	}

	return string(buf.Bytes())
}

func prepareFloat64(num string) float64 {
	f, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0.0
	}

	return f
}


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

func validateId(src string) string {
	if src == "0" {
		return ""
	}

	return src
}
