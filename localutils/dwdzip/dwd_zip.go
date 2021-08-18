package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	fileUrl := "https://opendata.dwd.de/climate_environment/CDC/observations_germany/climate/daily/kl/historical/tageswerte_KL_00001_19370101_19860630_hist.zip"
	fileName := "tageswerte_KL_00001_19370101_19860630_hist.zip";

	if err := DownloadFile(fileName, fileUrl); err != nil {
		panic(err)
	}
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = ReadZipFile(body)
	return err
}

func ReadZipFile(zipData []byte) error {
	r := bytes.NewReader(zipData)

	zipReader, err := zip.NewReader(r, r.Size())
	if err != nil {
		return err
	}


	for _, file := range zipReader.File {

		zippedFile, err := file.Open()
		if err != nil {
			return err
		}
		defer zippedFile.Close()

		if !file.FileInfo().IsDir() && isProductFileName(file.Name) {
			data, err := ioutil.ReadAll(zippedFile)
			if err != nil {
				return err
			}

			fmt.Println(file.Name)

			fmt.Println(string(data))
		} else {
			fmt.Println("!",file.Name)
		}
	}

	return nil
}

func isProductFileName(name string) bool {
	return strings.Index(name,"produkt_tu_stunde") > -1
}