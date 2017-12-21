package youtube

import (
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	URL     = "https://www.youtube.com"
	Results = URL + "/results" // results page
)

//Video :(Short for YoutubeVideo). This can be extended as wanted. (Views, uploader, etc)
type Video struct {
	URL   string
	Title string
}

//Search : Searches youtube for a video. Returns a slice of Video containing all search results.
func Search(query string) ([]Video, error) {

	var results []Video
	var vid Video

	u, err := url.Parse(Results)
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
				vid.URL = URL + href // current video (current search result)
				vid.Title = title
				results = append(results, vid) // append to total results
			}
		}

	})

	return results, nil
}

//Download : Downloads a youtube video and converts it into mp3 (using ffmpeg or avconv) Returns file path and err
func Download(video Video, Folder string) (string, error) {

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
