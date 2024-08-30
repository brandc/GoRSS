package RSSParse

import (
	"encoding/xml"
	"strings"
)

// Source: https://www.rssboard.org/rss-specification#ltcommentsgtSubelementOfLtitemgt

type Enclosure struct {
	Length int    `xml:"length,attr"`
	SMIME  string `xml:"smime,attr"`
	URL    string `xml:"url,attr"`
}

type Item struct {
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Title       string `xml:"title"`

	Author   string `xml:"author,omitempty"`
	Category string `xml:"category,omitempty"`
	Comment  string `xml:"comment,omitempty"`

	Enclosures []Enclosure `xml:"enclosure"`

	Guid            string `xml:"guid,omitempty"`
	PublicationDate string `xml:"pubDate,omitempty"`
}

type TextInput struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Name        string `xml:"name"`
	Link        string `xml:"link"`
}

type Image struct {
	Height int `xml:"height,attr"`
	Width  int `xml:"width,attr"`
	Title  int `xml:title,attr`

	Text string `xml:",chardata"`
	Href string `xml:"href,attr"`
}

type Link struct {
	Href string `xml:"href,attr"`
	URL  string `xml:",chardata"`
}

type Channel struct {
	// Required fields
	Description string `xml:"description"`
	Links       []Link `xml:"link"`
	Title       string `xml:"title"`

	Items []Item `xml:"item,omitempty"`

	// Optional fields
	Author          string      `xml:"author,omitempty"`
	Category        string      `xml:"category,,omitempty"`
	Cloud           string      `xml:"cloud,omitempty"`
	Copyright       string      `xml:"copyright,omitempty"`
	Documentation   string      `xml:"docs,omitempty"`
	Generator       string      `xml:"generator,omitempty"`
	Language        string      `xml:"language,omitempty"`
	LastBuildDate   string      `xml:"lastBuildDate,omitempty"`
	ManagingEditor  string      `xml:"managingEditor,omitempty"`
	PublicationDate string      `xml:"pubDate,omitempty"`
	Rating          string      `xml:"rating,omitempty"`
	SkipHours       string      `xml:"skipHours,omitempty"`
	SkipDays        string      `xml:"skipDays,omitempty"`
	TextInputs      []TextInput `xml:"textInput,omitempty"`
	TimeToLive      string      `xml:"ttl,omitempty"`
	WebMaster       string      `xml:"webMaster,omitempty"`

	// Fields with optional subelements
	Images     []Image `xml:"image,omitempty"`
	ImageWidth int     `xml:"image,attr"`
}

type RSS struct {
	XMLName  xml.Name
	Channels []Channel `xml:"channel"`
}

func RSSFeedParse(data string) (feed *RSS, err error) {
	decoder := xml.NewDecoder(strings.NewReader(data))
	err = decoder.Decode(&feed)
	if err != nil {
		return nil, err
	}

	return feed, nil
}
