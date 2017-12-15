package main

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

const (
	YTURL     = "https://www.youtube.com/"
	YTResults = YTURL + "results" // results page
)

// Searches youtube for a video. Returns a map[video title][video url]
func SearchYT(query string) (map[string]string, error) {
	m := make(map[string]string) // map of vid titles to vid urls

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

	doc.Find("a[rel=spf-prefetch]").Each(func(i int, s *goquery.Selection) {
		if title, ok := s.Attr("title"); ok {
			if href, ok := s.Attr("href"); ok {
				m[title] = YTURL + href
			}
		}

	})

	return m, nil
}
