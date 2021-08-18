package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("tr td:nth-of-type(7)", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.Visit("https://w1.weather.gov/data/obhistory/CYVR.html")
}