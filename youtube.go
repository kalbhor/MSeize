package main

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

const (
	YTURL     = "https://www.youtube.com"
	YTResults = YTURL + "/results" // results page
)

// Fields for a youtube video. This can be extended as wanted. (Views, uploader, etc)
type Video struct {
	URL   string
	Title string
}

// Searches youtube for a video. Returns a slice of Video containing all search results.
func SearchYT(query string) ([]Video, error) {

	var results []Video
	var vid Video

	u, err := url.Parse(YTResults)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Add("search_query", query)
	u.RawQuery = q.Encode() // Set url with query

	doc, err := goquery.NewDocument(u.String()) // parse page with results
	if err != nil {
		return nil, err
	}

	doc.Find("a[rel=spf-prefetch]").Each(func(i int, s *goquery.Selection) { // Go over all search results
		if title, ok := s.Attr("title"); ok {
			if href, ok := s.Attr("href"); ok {
				vid.URL = YTURL + href // current video (current search result)
				vid.Title = title
				results = append(results, vid) // append to total results
			}
		}

	})

	return results, nil
}
