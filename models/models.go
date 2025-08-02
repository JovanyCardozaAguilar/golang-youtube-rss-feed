package models

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"encoding/xml"
)
type Postgres struct {
	Db *pgxpool.Pool
}

type FeedProfile struct {
	VideoId string
	VChannelId string
	Title string
	Thumbnail string
	Watched bool
}

type ChannelProfile struct {
	ChannelId string
	Username string
	Avatar string
}

type VideoProfile struct {
	VideoId string
	VChannelId string
	Title string
	Thumbnail string
	Watched bool
}

type CategoryProfile struct {
	CategoryId string
	CatName string
}

type ChannelCategoryProfile struct {
	CcChannelId string
	CcCategoryId string
}

type VideoCategoryProfile struct {
	VcVideoId string
	VcCategoryId string
}

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
