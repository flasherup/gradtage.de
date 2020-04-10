package source

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"github.com/flasherup/gradtage.de/hrlaggregatorsvc/impl/parser"
	"github.com/go-kit/kit/log"
	"io/ioutil"
	"net/http"
	"strings"
)

type SourceDWD struct {
	url string
	logger log.Logger
}


func NewDWD(url string, logger log.Logger) *SourceDWD {
	return &SourceDWD{
		url: 	url,
		logger:	logger,
	}
}

func (sdwd SourceDWD) FetchTemperature(ch chan *parser.ParsedData, ids map[string]string) {
	for k,v := range ids {
		sdwd.fetchStation(k, v, ch)
	}
}

func (sdwd SourceDWD)fetchStation(id string, srcId string, ch chan *parser.ParsedData) {
	fileName := generateFileName(srcId)
	zip, err := DownloadFile(sdwd.url + fileName)
	if err != nil {
		ch <- &parser.ParsedData{ Success:false, StationID:id, Error:err }
		return
	}

	data, err := ReadZipFile(zip)
	if err != nil {
		ch <- &parser.ParsedData{ Success:false, StationID:id, Error:err }
		return
	}

	temps, err := parser.ParseDWD(data)
	if err != nil {
		ch <- &parser.ParsedData{ Success:false, StationID:id, Error:err }
		return
	}

	res := parser.ParsedData{
		Success:true,
		StationID:id,
		Temps:*temps,
	}

	ch <- &res
}

func DownloadFile(url string) (*[]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &body, err
}

func ReadZipFile(zipData *[]byte) (*[]byte, error) {
	r := bytes.NewReader(*zipData)

	zipReader, err := zip.NewReader(r, r.Size())
	if err != nil {
		return nil,err
	}

	for _, file := range zipReader.File {
		zippedFile, err := file.Open()
		if err != nil {
			return nil,err
		}
		defer zippedFile.Close()

		if !file.FileInfo().IsDir() && isProductFileName(file.Name) {
			data, err := ioutil.ReadAll(zippedFile)
			if err != nil {
				return nil,err
			}
			return &data, nil
		}
	}

	return nil,errors.New("document file not find")
}

func isProductFileName(name string) bool {
	return strings.Index(name,"produkt_tu_stunde") > -1
}

func generateFileName(id string) string {
	return fmt.Sprintf("stundenwerte_TU_%s_akt.zip", id)
}