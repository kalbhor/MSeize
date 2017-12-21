package main

import (
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	youtubeURL     = "https://www.youtube.com"
	youtubeResults = youtubeURL + "/results" // results page
)

//YTVideo :(Short for YoutubeVideo). This can be extended as wanted. (Views, uploader, etc)
type YTVideo struct {
	URL   string
	Title string
}

//SearchYT : Searches youtube for a video. Returns a slice of Video containing all search results.
func SearchYT(query string) ([]YTVideo, error) {

	var results []YTVideo
	var vid YTVideo

	u, err := url.Parse(youtubeResults)
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
				vid.URL = youtubeURL + href // current video (current search result)
				vid.Title = title
				results = append(results, vid) // append to total results
			}
		}

	})

	return results, nil
}

//DownloadMP3 : Downloads a youtube video and converts it into mp3 (using ffmpeg or avconv) Returns file path and err
func DownloadMP3(video YTVideo, Folder string) (string, error) {

	// Download path and filePath should essentially be the same thing,
	// but passing filePath in youtube-dl causes an extension error
	// since the file goes from webm -> .mp3,
	// having .mp3 in the filepath causes problems in the first step (downloading webm)
	filePath := path.Join(Folder, strings.Replace(video.Title, "\"", "'", -1)+".mp3")
	downloadPath := path.Join(Folder, "%(title)s.%(ext)s")

	cmd := exec.Command("youtube-dl", "--extract-audio", "--output", downloadPath,
		"--audio-format", "mp3", video.URL) // youtube-dl command

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	return filePath, nil
}
