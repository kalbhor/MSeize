package main

import (
	"testing"
)

var (
	VideoTitle = "Charlie bit my finger - again !"
	VideoURL   = "https://www.youtube.com/watch?v=_OBlgSz8sSM"
	query      = "charlie bit my finger"
)

func TestSearchYT(t *testing.T) {
	results, err := SearchYT(query)
	check := false

	if err != nil {
		t.Fatalf("Could not scrape YT for query. %v", err)
	}

	if len(results) < 1 {
		t.Fatalf("Results are empty")
	}

	for _, result := range results {
		if result.Title == VideoTitle && result.URL == VideoURL {
			check = true
			break
		}

	}

	if !check {
		t.Fatalf("Correct results were not found")
	}

}
