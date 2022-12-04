package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber"
	"regexp"
)

func main() {
	app := fiber.New()

	// Define the route for the tool
	app.Get("/links", func(c *fiber.Ctx) {
		// Get the URL from the query parameters
		url := c.Query("url")

		// Use goquery to scrape the page for links
		doc, err := goquery.NewDocument(url)
		if err != nil {
			c.Status(500).Send(err)
			return
		}

		// Use a regular expression to extract the links
		var links []string
		re := regexp.MustCompile(`https?:\/\/[^\s]+`)
		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			href, _ := s.Attr("href")
			links = append(links, re.FindString(href))
		})

		// Return the links as JSON in the HTTP response
		c.JSON(links)
	})

	// Start the server on port 3000
	app.Listen(3000)
}
