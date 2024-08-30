package rss

import (
	"bytes"
	"encoding/xml"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var comprss = Rss{
	XMLName:     xml.Name{Space: "", Local: "rss"},
	Version:     "2.0",
	Title:       "Testfeed",
	Description: "This is a Testfeed for golang-rss-rw",
	Link:        "http://example.com/test.rss",

	Item: []Item{
		{
			Title:       "foo",
			Link:        "http://example.com/foo",
			Description: "lorem ipsum",
			PubDate:     "Sun, 09 Aug 2015 16:05:14 +0200",
		},
		{
			Title:       "bar",
			Link:        "http://example.com/bar",
			Description: "dolor sit amet",
			PubDate:     "Sun, 09 Aug 2015 16:05:14 +0200",
		},
	},
}

func TestParseRSS(t *testing.T) {
	file, err := os.Open("result.xml")
	if err != nil {
		t.Errorf("Failed to open result.xml: %v", err)
	}

	rss, err := ParseRSS(file)
	if err != nil {
		t.Errorf("ParseRSS returned an error: %v", err)
	}

	assert.Equal(t, comprss, rss)

}

func TestWriteRSS(t *testing.T) {
	file, err := os.Create("testout.rss")
	if err != nil {
		t.Errorf("Failed to create testout.rss: %v", err)
	}

	err = comprss.WriteRSS(file)
	if err != nil {
		t.Errorf("WriteRSS returned an error: %v", err)
	}

	file.Close()

	// now test equality

	file1, err := os.Open("result.xml")
	if err != nil {
		t.Errorf("Failed to open result.xml: %v", err)
	}
	file2, err := os.Open("testout.rss")
	if err != nil {
		t.Errorf("Failed to open testout.rss: %v", err)
	}

	// compare the two files
	buf1 := new(bytes.Buffer)
	buf2 := new(bytes.Buffer)

	buf1.ReadFrom(file1)
	buf2.ReadFrom(file2)

	assert.Equal(t, buf1.String(), buf2.String())

	file1.Close()
	file2.Close()

	// cleanup
	os.Remove("testout.rss")
}
