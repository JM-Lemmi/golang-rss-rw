package rss

import (
	"encoding/xml"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", rssHandler)
	http.ListenAndServe(":3000", nil)
}

func rssHandler(w http.ResponseWriter, r *http.Request) {
	type Item struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		PubDate     string `xml:"pubDate"`
	}

	type rss struct {
		Version     string `xml:"version,attr"`
		Description string `xml:"channel>description"`
		Link        string `xml:"channel>link"`
		Title       string `xml:"channel>title"`

		Item []Item `xml:"channel>item"`
	}

	articles := []Item{
		{"foo", "http://mywebsite.com/foo", "lorem ipsum", time.Now().Format(time.RFC1123Z)},
		{"foo2", "http://mywebsite.com/foo2", "lorem ipsum2", time.Now().Format(time.RFC1123Z)}}

	feed := &rss{
		Version:     "2.0",
		Description: "My super website",
		Link:        "http://mywebsite.com",
		Title:       "Mywebsite",
		Item:        articles,
	}

	x, err := xml.MarshalIndent(feed, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}
