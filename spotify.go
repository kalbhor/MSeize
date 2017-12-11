package main

import (
	"context"
	"errors"
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
	Image       string
	DiscNumber  int
	TrackNumber int
}

//Sets values from search results
func (m *Metadata) Load(track spotify.FullTrack) {
	m.Title = track.SimpleTrack.Name
	m.Artist = track.SimpleTrack.Artists[0].Name
	m.Album = track.Album.Name
	m.Image = track.Album.Images[0].URL
	m.DiscNumber = track.SimpleTrack.DiscNumber
	m.TrackNumber = track.SimpleTrack.TrackNumber
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

	m.Load(results.Tracks.Tracks[0])
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
