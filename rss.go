package rss

import (
	"encoding/xml"
	"io"
)

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type Rss struct {
	XMLName     xml.Name `xml:"rss"`
	Version     string   `xml:"version,attr"`
	Title       string   `xml:"channel>title"`
	Description string   `xml:"channel>description"`
	Link        string   `xml:"channel>link"`

	Item []Item `xml:"channel>item"`
}

func ParseRSS(r io.Reader) (Rss, error) {
	var rss Rss
	err := xml.NewDecoder(r).Decode(&rss)
	if err != nil {
		return Rss{}, err
	}
	return rss, nil
}

func (rss Rss) WriteRSS(w io.Writer) error {
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	return enc.Encode(rss)
}
