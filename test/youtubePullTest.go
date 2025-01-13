package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Entries []Entry  `xml:"entry"`
}

type Entry struct {
	Text      string `xml:",chardata"`
	VideoId   string `xml:"videoId"`
	ChannelId string `xml:"channelId"`
	Title     string `xml:"title"`
	Author    struct {
		Text string `xml:",chardata"`
		Name string `xml:"name"`
		URI  string `xml:"uri"`
	} `xml:"author"`
	Group struct {
		Text    string `xml:",chardata"`
		Title   string `xml:"title"`
		Content struct {
			Text   string `xml:",chardata"`
			URL    string `xml:"url,attr"`
			Type   string `xml:"type,attr"`
			Width  string `xml:"width,attr"`
			Height string `xml:"height,attr"`
		} `xml:"content"`
		Thumbnail struct {
			Text   string `xml:",chardata"`
			URL    string `xml:"url,attr"`
			Width  string `xml:"width,attr"`
			Height string `xml:"height,attr"`
		} `xml:"thumbnail"`
	} `xml:"group"`
}

/*
func (e Entry) string() {
	return fmt.Sprintf(e.ChannelId)
}
*/

func main() {
	resp, _ := http.Get("https://www.youtube.com/feeds/videos.xml?channel_id=UCq6VFHwMzcMXbuKyG7SQYIg")
	bytes, _ := io.ReadAll(resp.Body)
	//string_body := string(bytes)
	//fmt.Println(string_body)
	resp.Body.Close()

	var f Feed
	xml.Unmarshal(bytes, &f)

	fmt.Println(f.Entries[0])

}
