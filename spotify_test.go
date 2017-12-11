package main

import "testing"

func TestGetMetadata(t *testing.T) {
	client, err := Auth()
	if err != nil {
		t.Fatalf("Could not authorise. %v", err)
	}
	m, err := GetMetadata("Comfortably Numb", client)
	if err != nil {
		t.Fatalf("Could not search spotify. %v", err)
	}
	if m.Title != "Comfortably Numb" {
		t.Errorf("Got wrong title %s", m.Title)
	}
	if m.Artist != "Pink Floyd" {
		t.Errorf("Got wrong artist %s", m.Artist)
	}
	if m.Album != "The Wall" {
		t.Errorf("Got wrong album %s", m.Album)
	}
	if m.DiscNumber != 2 {
		t.Errorf("Got wrong disc number %v", m.DiscNumber)
	}
	if m.TrackNumber != 6 {
		t.Errorf("Got wrong track number %v", m.TrackNumber)
	}
}
