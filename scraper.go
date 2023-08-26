package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Industry struct {
	Url, Image, Name string
}

func main() {
	fmt.Println("Enter the website you wish to scrape")
	var scrapeURL string
	fmt.Scanln(&scrapeURL)

	// set collector
	c := colly.NewCollector()

	// set user-agent header
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	// scrape logic
	c.Visit(scrapeURL)

	var industries []Industry
	// iterating over the list of industry card

	// HTML elements

	c.OnHTML(".elementor-element-6b05593c .section_cases__item", func(e *colly.HTMLElement) {

		url := e.Attr("href")
		image := e.ChildAttr(".elementor-image-box-img img", "data-lazy-src")
		name := e.ChildText(".elementor-image-box-content .elementor-image-box-title")

		// filter out unwanted data
		if url != "" || image != "" || name != "" {
			// initialize a new Industry instance
			industry := Industry{
				Url:   url,
				Image: image,
				Name:  name,
			}
			// add the industry instance to the list
			// of scraped industries
			industries = append(industries, industry)
		}
	})

	file, err := os.Create("industries.json")

	if err != nil {
		log.Fatal("Failed to create output JSON file:", err)
	}
	defer file.Close()

	// convert industries to an indented JSON string
	jsonString, _ := json.MarshalIndent(industries, " ", " ")

	// write the JSON string to the file
	file.Write(jsonString)
}
