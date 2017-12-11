package main

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/zmb3/spotify"
)

//Structure for one track's metadata
type Metadata struct {
	Client      spotify.Client
	Title       string
	Artist      string
	Album       string
	Image       []byte
	DiscNumber  int
	TrackNumber int
}

//Sets values from search results
func (m *Metadata) Load(track spotify.FullTrack) error {
	m.Title = track.SimpleTrack.Name
	m.Artist = track.SimpleTrack.Artists[0].Name
	m.Album = track.Album.Name
	m.DiscNumber = track.SimpleTrack.DiscNumber
	m.TrackNumber = track.SimpleTrack.TrackNumber
	imageURL := track.Album.Images[0].URL

	resp, err := http.Get(imageURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	m.Image = b

	return nil
}

// Searches spotify and returns a loaded metadata struct
func GetMetadata(query string, client spotify.Client) (*Metadata, error) {

	m := new(Metadata)

	results, err := client.Search(query, spotify.SearchTypeTrack)
	if err != nil {
		return nil, err
	} else if len(results.Tracks.Tracks) == 0 { // Search results were empty
		return nil, errors.New("Couldn't fetch metadata")
	}

	err = m.Load(results.Tracks.Tracks[0]) // Pass in the top result
	if err != nil {
		return m, err
	}
	return m, nil

}

//Returns a usable "client" that can request spotify content
func Auth() (spotify.Client, error) {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		return spotify.Authenticator{}.NewClient(&oauth2.Token{}), err
	}

	client := spotify.Authenticator{}.NewClient(token)

	return client, nil
}
