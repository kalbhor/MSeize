package main

import (
	"log"
	"os"
	"path"
	"testing"
)

var (
	VideoTitle = "Charlie bit my finger - again !"
	VideoURL   = "https://www.youtube.com/watch?v=_OBlgSz8sSM"
	query      = "charlie bit my finger"
)

func TestYT(t *testing.T) {
	results, err := SearchYT(query)
	var video YTVideo
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
			video = result
			break
		}

	}

	if !check {
		t.Fatalf("Correct results were not found")
	}

	filePath, err := DownloadMP3(video, "./tests_files")
	if err != nil {
		t.Fatalf("Could not download mp3 for %v", video.URL)
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Fatalf("Downloaded file couldn't be found. %v", filePath)
	}

	// Tests have passed, clean up time
	err = os.RemoveAll(path.Dir(filePath))
	if err != nil {
		log.Fatalf("%v", err)

	}

}
