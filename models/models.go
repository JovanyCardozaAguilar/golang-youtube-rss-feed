package models

import "github.com/jackc/pgx/v5/pgxpool"

type Postgres struct {
	Db *pgxpool.Pool
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
