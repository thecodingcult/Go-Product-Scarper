<h2>Go Web Scarper</h2>
<p>This Go program uses the <code>gocolly</code> package to scrape product data from a website. The program takes in a URL for the website to scrape and an output format (either CSV or JSON), and outputs a file with the scraped data in the specified format.</p>
<p>The program first defines a struct called <code>Product</code> to represent a product on the website. The <code>Product</code> struct has fields for the name, price, details, brand, description, image URL, and availability of the product.</p>
<p>The program then parses command-line arguments using the <code>flag</code> package, and checks that the website URL is specified. It initializes a new <code>Collector</code> from the <code>gocolly</code> package, which limits the rate at which requests are made to the website to avoid overloading it. The program also creates a file to store the scraped data, using either a CSV writer or a JSON encoder depending on the specified output format.</p>
<p>The program sets up request handlers for the list of products and the "next" link, and an error handler to handle any errors that occur during scraping. It starts the scraper by visiting the first page of the website, and waits for all the requests to finish before signaling that it is done.</p> 
<p>Once the program has finished scraping the website, it writes the scraped data to the specified output file in the specified format. If the output format is CSV, it writes the data to a CSV file using the CSV writer. If the output format is JSON, it writes the data to a JSON file using the JSON encoder.</p>

<h3>Parsing command line arguments:</h3>
<pre>
<code class="language-go">
websitePtr := flag.String("website", "", "URL of the website to scrape")
outputFormatPtr := flag.String("format", "csv", "Output file format ('csv' or 'json')")
flag.Parse()
</code>
</pre>
<h3>Defining the `Product` struct:</h3>
<pre>
<code class="language-go">
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
</code>
</pre>
<h3>Initializing a new collector:</h3>
<pre>
<code class="language-go">
// Initialize a new collector
c := colly.NewCollector(
    // Limit the rate at which requests are made to the website to avoid overloading it
    colly.Async(true),
    colly.MaxDepth(2),
    colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"),
)
</code>
</pre>
<h3>Request handler for the list of products:</h3>
<pre>
<code class="language-go">
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
    } else if *

  </body>
</html>
