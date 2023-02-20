package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/gocolly/colly"
)

// Product represents a product on the website
type Product struct {
	Name        string `json:"name"`
	Price       string `json:"price"`
	Details     string `json:"details"`
	Brand       string `json:"brand"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Availability string `json:"availability"`
}

func main() {
	// Parse command-line arguments
	websitePtr := flag.String("website", "", "URL of the website to scrape")
	outputFormatPtr := flag.String("format", "csv", "Output file format ('csv' or 'json')")
	flag.Parse()

	// Check that the website URL is specified
	if *websitePtr == "" {
		log.Fatal("Error: website URL is required")
	}

	// Initialize a new collector
	c := colly.NewCollector(
		// Limit the rate at which requests are made to the website to avoid overloading it
		colly.Async(true),
		colly.MaxDepth(2),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"),
	)

	// Create a file to store the scraped data
	var file *os.File
	var writer *csv.Writer
	var encoder *json.Encoder
	var err error
	if *outputFormatPtr == "csv" {
		file, err = os.Create("products.csv")
		if err != nil {
			log.Fatal("Error creating file:", err)
		}
		defer file.Close()

		// Create a CSV writer to write data to the file
		writer = csv.NewWriter(file)
		defer writer.Flush()
	} else if *outputFormatPtr == "json" {
		file, err = os.Create("products.json")
		if err != nil {
			log.Fatal("Error creating file:", err)
		}
		defer file.Close()

		// Create a JSON encoder to write data to the file
		encoder = json.NewEncoder(file)
		encoder.SetIndent("", "  ")
	} else {
		log.Fatal("Error: invalid output file format. Must be 'csv' or 'json'")
	}

	// Create a channel to signal when all the requests are done
	done := make(chan bool)

	// Set up the request handler for the list of products
	c.OnHTML("div.product-item", func(e *colly.HTMLElement) {
		product := Product{
			Name:    e.ChildText("a.product-title-link"),
			Price:   e.ChildText("span.product-price"),
			Details: e.ChildText("div.product-details"),
		}

		if *outputFormatPtr == "csv" {
			// Write the data to the CSV file
			if err := writer.Write([]string{product.Name, product.Price, product.Details}); err != nil {
				log.Println("Error writing record to CSV:", err)
			}
		} else if *outputFormatPtr == "json" {
			// Write the data to the JSON file
			if err := encoder.Encode(product); err != nil {
				log.Println("Error writing record to JSON:", err)
			}
		}
	})

	// Set up the request handler for the "next" link
	c.OnHTML("li.next a", func(e *colly.HTMLElement) {

			// Visit the next page
			link := e.Attr("href")
			if err := e.Request.Visit(link); err != nil {
				log.Println("Error visiting next page:", err)
			}
		})
	
		// Set up the error handler
		c.OnError(func(r *colly.Response, err error) {
			log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		})
	
		// Start the scraper by visiting the first page
		c.Visit(*websitePtr)
	
		// Wait for all the requests to finish
		c.Wait()
	
		// Signal that we are done
		done <- true
	}
	
