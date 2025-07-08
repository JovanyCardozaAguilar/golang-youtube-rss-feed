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
	Title string
	Thumbnail string
	Watched bool
	VideoChannel string
}

type CategoryProfile struct {
	CategoryId string
	CatName string
	CatChannel string
}
