package main

import (
	"fmt"
	"github.com/flasherup/gradtage.de/utils/weatherbithistorical/parser"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	err := getPureWeatherbitData("KDCA","2023-02-01", "2023-02-08")
	if err != nil {
		log.Fatal(err)
	}
}

func getPureWeatherbitData(station, start, end string) error {
	key := "955ab65523e4475da39681fba22f8b13"
	//key := "6ea2586276b5445ea34614d8e45596de"
	url := fmt.Sprintf("https://api.weatherbit.io/v2.0/history/daily?station=%s&key=%s&start_date=%s&end_date=%s", station, key, start, end)
	fmt.Println("msg", "weather bit request", "url", url)

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(contents))

	result, err := parser.ParseWeatherBit(&contents)
	if err != nil {
		return err
	}

	fmt.Println(result)
	return nil
}
