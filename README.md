# Bhailink

Bhailink is a tool for scraping and extracting links from web pages. It allows you to specify a URL and depth, and it will crawl the page and return all of the links it finds.

## Features

- Scrape links from a single page or multiple pages, depending on the depth specified
- Filter the links that are returned, based on patterns or exclusion rules
- Persist the links in a database or file for later use

## Usage

To use Bhailink, you can send a GET request to the `/links` route, with the following query parameters:

- `url`: The URL of the page to scrape.
- `depth`: The depth of the scraping, i.e. how many levels of pages to follow the links on.

For example, to scrape the links on the page at `https://example.com` to a depth of 1, you can send a request like this:

$GET /links?url=https://example.com&depth=1


The response will be a JSON array of the links found on the page.

## Installation

To install Bhailink, clone the repository and run `go build` to build the binary. Then, you can run the binary to start the server and use Bhailink.

Alternatively, you can use `go get` to install Bhailink and its dependencies:

$go get github.com/hangyakuzero/bhailink


## License

Bhailink is released under the MIT License. See the [LICENSE](LICENSE) file for more details.

