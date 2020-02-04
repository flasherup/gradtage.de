package parser

import (
	"bytes"
	"encoding/csv"
	"github.com/flasherup/gradtage.de/common"
	"github.com/flasherup/gradtage.de/hourlysvc"
	"io"
	"strconv"
	"strings"
	"time"
)

func ParseNOAA(data *[]byte) (*[]hourlysvc.Temperature, error) {
	r := bytes.NewReader(*data)
	reader := csv.NewReader(r)
	reader.Comma = ';'
	temps := make([]hourlysvc.Temperature, 0)
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
		temp := line[3]

		//2006-01-02T15:04:05Z
		pDate, err := time.Parse("2006010215", date)
		if err != nil {
			return nil, err
		}
		dValue := pDate.Format(common.TimeLayout)

		temp = strings.TrimSpace(temp)
		value, err := strconv.ParseFloat(temp, 64)
		if err != nil {
			return nil, err
		}

		temps = append(temps, hourlysvc.Temperature{
			Date:    dValue,
			Temperature:   value,
		})

	}
	return &temps,nil
}