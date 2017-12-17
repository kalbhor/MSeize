package main

import (
	"net/url"
	"os/exec"
	"strconv"
	"time"

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

//MusicFile : Represents a downloaded mp3 file
type MusicFile struct {
	FolderPath string // unique unix epoch (unique to a nano second)
	Path       string // path of the music file
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
func DownloadMP3(video YTVideo) (*MusicFile, error) {

	file := new(MusicFile)

	epoch := time.Now().UnixNano() // This will fail after the year 2262
	epochString := strconv.FormatInt(epoch, 10)

	file.FolderPath = "./temp/" + epochString + `/`
	file.Path = file.FolderPath + video.Title + ".mp3"
	downloadPath := file.FolderPath + "%(title)s.%(ext)s"
	// Download path and file.Path should essentially be the same thing
	// but having ".mp3" in the file.Path causes problems with youtube-dl
	// The downloadPath variable uses a much more flexible approach and uses the video title

	cmd := exec.Command("youtube-dl", "--extract-audio", "--output", downloadPath,
		"--audio-format", "mp3", video.URL) // youtube-dl command

	cmd.Run()
	cmd.Wait()
	/*stdout, err := cmd.StdoutPipe() // stderr logging
	if err != nil {
		return filePath, err
	}
	if err := cmd.Start(); err != nil {
		return filePath, err
	}

	// logging youtube-dl command
	in := bufio.NewScanner(stdout)
	for in.Scan() {
		log.Printf(in.Text())
	}
	if err := in.Err(); err != nil {
		log.Printf("error : %s", err)
	}*/

	return file, nil
}
