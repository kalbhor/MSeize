package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bogem/id3v2"
	"github.com/kalbhor/MSeize/spotify"
	"github.com/kalbhor/MSeize/youtube"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var query string
	var choice int

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter song name : ")
	query, _ = reader.ReadString('\n')
	query = strings.TrimSuffix(query, "\n")

	videos, err := youtube.Search(query)
	checkErr(err)

	for i, video := range videos {
		fmt.Printf("[%v] %v\n", i+1, video.Title)
	}
	fmt.Print("Enter song number : ")
	fmt.Scan(&choice)

	if choice > len(videos) {
		log.Fatal("Wrong input")
	}

	fmt.Println("--------------")

	path, err := youtube.Download(videos[choice-1], "./")
	checkErr(err)

	client, err := spotify.Auth()
	checkErr(err)

	metadata, err := spotify.GetMetadata(client, query)
	checkErr(err)

	tag, err := id3v2.Open(path, id3v2.Options{Parse: true})
	checkErr(err)
	defer tag.Close()

	tag.SetTitle(metadata.Title)
	tag.SetAlbum(metadata.Album)

	artists := strings.Join(metadata.Artists, ",")
	tag.SetArtist(artists)

	pic := id3v2.PictureFrame{
		Encoding:    id3v2.EncodingUTF8,
		MimeType:    "image/jpeg",
		PictureType: id3v2.PTFrontCover,
		Description: "Front cover",
		Picture:     metadata.Image,
	}

	tag.AddAttachedPicture(pic)

	TrackNumber := strconv.Itoa(metadata.TrackNumber)
	trackFrame := id3v2.TextFrame{
		Encoding: id3v2.EncodingUTF8,
		Text:     TrackNumber,
	}

	tag.AddFrame("TRCK", trackFrame)

	DiscNumber := strconv.Itoa(metadata.DiscNumber)
	discFrame := id3v2.TextFrame{
		Encoding: id3v2.EncodingUTF8,
		Text:     DiscNumber,
	}

	tag.AddFrame("TPOS", discFrame)

	if err = tag.Save(); err != nil {
		log.Fatal("Error : ", err)
	}

}
