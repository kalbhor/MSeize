package main

import "testing"

var (
	Title       = "Riding With The King"
	Artists     = []string{"Eric Clapton", "B.B. King"}
	Album       = "Riding With The King"
	DiscNumber  = 1
	TrackNumber = 1
)

func TestGetMetadata(t *testing.T) {
	client, err := SpotifyAuth()
	if err != nil {
		t.Fatalf("Could not authorise. %v", err)
	}

	m, err := GetMetadata("Riding with the king", client)
	if err != nil {
		t.Fatalf("Could not search spotify. %v", err)
	}

	if m.Title != Title {
		t.Fatalf("Didn't get correct title. Got %s instead", m.Title)
	}
	if m.Album != Album {
		t.Fatalf("Didn't get correct Album. Got %s instead", m.Album)
	}
	if m.TrackNumber != TrackNumber {
		t.Fatalf("Didn't get correct track number. Got %v instead", m.TrackNumber)
	}
	if m.DiscNumber != DiscNumber {
		t.Fatalf("Didn't get correct disc number. Got %v instead", m.DiscNumber)
	}
	if m.Artists[0] != "Eric Clapton" || m.Artists[1] != "B.B. King" {
		t.Fatalf("Didn't get correct artits")
	}

}
