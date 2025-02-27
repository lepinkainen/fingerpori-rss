package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/feeds"
)

// FingerporiItem represents a single comic item from the JSON API
type FingerporiItem struct {
	ID          int64  `json:"id"`
	Href        string `json:"href"`
	DisplayDate string `json:"displayDate"`
	Title       string `json:"title"`
	Picture     struct {
		ID           int64  `json:"id"`
		Width        int    `json:"width"`
		Height       int    `json:"height"`
		URL          string `json:"url"`
		SquareURL    string `json:"squareUrl"`
		Photographer string `json:"photographer"`
	} `json:"picture"`
	PaidType       string                 `json:"paidType"`
	Category       string                 `json:"category"`
	SectionTheme   string                 `json:"sectionTheme"`
	InfoRowContent map[string]interface{} `json:"infoRowContent"`
	Tags           []string               `json:"tags"`
}

func main() {
	// Define command-line flags
	outDir := flag.String("outdir", ".", "Directory where the RSS feed XML file will be created")
	flag.Parse()

	// URL to fetch the JSON data from
	url := "https://www.hs.fi/api/laneitems/39221/list/normal/290"

	// Fetch the JSON data
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Parse the JSON data
	var items []FingerporiItem
	err = json.Unmarshal(body, &items)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// Create a new feed
	now := time.Now()
	feed := &feeds.Feed{
		Title:       "Fingerpori Comics",
		Link:        &feeds.Link{Href: "https://www.hs.fi/fingerpori/"},
		Description: "Daily Fingerpori comics from Helsingin Sanomat",
		Author:      &feeds.Author{Name: "Pertti Jarla"},
		Created:     now,
	}

	// Add items to the feed
	feed.Items = make([]*feeds.Item, 0, len(items))
	for _, item := range items {
		// Parse the display date
		displayDate, err := time.Parse("2006-01-02T15:04:05.000-07:00", item.DisplayDate)
		if err != nil {
			log.Printf("Error parsing date %s: %v", item.DisplayDate, err)
			displayDate = now
		}

		// Create a direct image URL with replacements
		baseURL := "https://images.sanoma-sndp.fi/"
		imageID := strings.Split(item.Picture.URL, "/")[3]
		imageURL := fmt.Sprintf("%s%s/normal/1440.jpg", baseURL, imageID)

		// Create a simple content string with the image
		content := fmt.Sprintf("<img src=\"%s\" alt=\"%s\">", imageURL, item.Title)

		// Create the feed item
		feedItem := &feeds.Item{
			Id:          fmt.Sprintf("https://www.hs.fi%s", item.Href),
			Title:       fmt.Sprintf("%s - %s", item.Title, displayDate.Format("2006-01-02")),
			Link:        &feeds.Link{Href: fmt.Sprintf("https://www.hs.fi%s", item.Href)},
			Description: content,
			Created:     displayDate,
			Author:      &feeds.Author{Name: item.Picture.Photographer},
		}

		feed.Items = append(feed.Items, feedItem)
	}

	// Generate the RSS feed
	rss, err := feed.ToRss()
	if err != nil {
		log.Fatalf("Error generating RSS: %v", err)
	}

	// Ensure the output directory exists
	if err := os.MkdirAll(*outDir, 0755); err != nil {
		log.Fatalf("Error creating output directory: %v", err)
	}

	// Create the output file path
	outputPath := filepath.Join(*outDir, "fingerpori.xml")

	// Write to the output file
	file, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(rss)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	log.Printf("RSS feed generated successfully and saved to %s", outputPath)
}
