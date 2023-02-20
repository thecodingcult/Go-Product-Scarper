<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
</head>
<body>
  <h1>My Scraper</h1>

  <p>This is a command-line scraper that extracts information about products from a website.</p>

  <h2>Usage</h2>

  <p>To use the scraper, run the following command:</p>

  <pre><code>go run main.go -website <i>website_url</i> -format <i>output_format</i></code></pre>

  <p>Where <code><i>website_url</i></code> is the URL of the website to scrape and <code><i>output_format</i></code> is the desired output file format ('csv' or 'json').</p>

  <h2>Dependencies</h2>

  <ul>
    <li><a href="https://github.com/gocolly/colly">colly</a></li>
  </ul>

  <h2>License</h2>

  <p>This project is licensed under the MIT License. See the <a href="LICENSE">LICENSE</a> file for details.</p>
</body>
</html>
