package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Source: https://www.rssboard.org/rss-specification#ltcommentsgtSubelementOfLtitemgt

type Enclosure struct {
	Length int `xml:"length"`
	SMIME string `xml:"smime"`
	URL string `xml:"url"`
}

type Item struct {
	Description string `xml:"description"`
	Link string `xml:"link"`
	Title string `xml:"title"`

	Author string `xml:"author,omitempty"`
	Category string `xml:"category,omitempty"`
	Comment string `xml:"comment,omitempty"`

	Enclosures []Enclosure `xml:"enclosure,omitempty"`

	Guid string `xml:"guid,omitempty"`
	PublicationDate string `xml:"pubDate,omitempty"`
}

type TextInput struct {
	Title string `xml:"title"`
	Description string `xml:"description"`
	Name string `xml:"name"`
	Link string `xml:"link"`
}

type Channel struct {
	// Required fields
	Description string `xml:"description"`
	Link string `xml:"link"`
	Title string `xml:"title"`

	Items []Item `xml:"item,omitempty"`

	// Optional fields
	Author string `xml:"author,omitempty"`
	Category string `xml:"category,,omitempty"`
	Cloud string `xml:"cloud,omitempty"`
	Copyright string `xml:"copyright,omitempty"`
	Documentation string `xml:"docs,omitempty"`
	Generator string `xml:"generator,omitempty"`
	Image string `xml:"image,omitempty"`
	Language string `xml:"language,omitempty"`
	LastBuildDate string `xml:"lastBuildDate,omitempty"`
	ManagingEditor string `xml:"managingEditor,omitempty"`
	PublicationDate string `xml:"pubDate,omitempty"`
	Rating string `xml:"rating,omitempty"`
	SkipHours string `xml:"skipHours,omitempty"`
	SkipDays string `xml:"skipDays,omitempty"`
	TextInputs []TextInput `xml:"textInput,omitempty"`
	TimeToLive string `xml:"ttl,omitempty"`
	WebMaster string `xml:"webMaster,omitempty"`
}

func RSSFeedParse(feed string) {
	return
}
















