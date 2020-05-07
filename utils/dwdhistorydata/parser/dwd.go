package parser

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/dailysvc"
	"io"
	"strconv"
	"strings"
	"time"
)

type ParsedData struct {
	Success bool
	StationID string
	Temps []dailysvc.Temperature
	Error error
}

func ParseDWD(data *[]byte) (*[]dailysvc.Temperature, error) {
	r := bytes.NewReader(*data)
	reader := csv.NewReader(r)
	reader.Comma = ';'
	temps := make([]dailysvc.Temperature, 0)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err

		}
		if line[0] == "STATIONS_ID" {
			continue
		}

		date := line[1]
		temp := line[13]

		if temp == "-999" {
			fmt.Println("Skip data", date)
			continue
		}

		//2006-01-02T15:04:05Z
		pDate, err := time.Parse("20060102", date)
		if err != nil {
			return nil, err
		}
		dValue := pDate.Format(common.TimeLayout)

		temp = strings.TrimSpace(temp)
		value, err := strconv.ParseFloat(temp, 64)
		if err != nil {
			return nil, err
		}

		temps = append(temps, dailysvc.Temperature{
			Date:    dValue,
			Temperature:   value,
		})

	}
	return &temps,nil
}