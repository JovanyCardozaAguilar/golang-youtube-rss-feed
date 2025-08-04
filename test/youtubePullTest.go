package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"regexp"
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
	Published string `xml:"published"`
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

func retrieveChannelID(username string) (string) {
	url := "https://www.youtube.com/@" + username
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	req.Header.Set("Accept", "text/html")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Get channelID
	re := regexp.MustCompile(`"browseId":"([0-9A-Za-z_-]{24})"`)
	match := re.FindSubmatch(bytes)
	fmt.Printf("matches (%d): %s", len(match), match)
	if len(match) < 2 {
		fmt.Println("DEBUG:")
		fmt.Println(string(bytes)) 
		panic("channelId (browseId) not found in page source")
	}

	return string(match[1])
}

func retrieveAvatarURL(username string) (string) {
	url := "https://www.youtube.com/@" + username

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")
	req.Header.Set("Accept", "text/html")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`<meta\s+property=["']og:image["']\s+content=["'](https://[^"']*yt3\.[^"']+)["']`)
	match := re.FindSubmatch(body)
	if len(match) < 2 {
		panic("avatar image not found in page")
	}

	return string(match[1])
}


func main() {
	url := fmt.Sprintf("https://www.youtube.com/feeds/videos.xml?channel_id=%s", retrieveChannelID("Acerola_t"))
	resp, _ := http.Get(url)
	bytes, _ := io.ReadAll(resp.Body)
	//string_body := string(bytes)
	//fmt.Println(string_body)
	resp.Body.Close()

	avatar := retrieveAvatarURL("Acerola_t")

	var f Feed
	xml.Unmarshal(bytes, &f)

	fmt.Println("\nFull Entry Test:")
	fmt.Println(f.Entries[0])
	fmt.Println("\nVideos:")
	for _, entry := range f.Entries {
		fmt.Printf("ChannelID: %s, Youtuber: %s, Avatar: %s, VideoID: %s, Title: %s, Published: %s, Thumbnail: %s \n", entry.ChannelId, entry.Author.Name, avatar, entry.VideoId, entry.Title, entry.Published, entry.Group.Thumbnail.URL)
	}

}
