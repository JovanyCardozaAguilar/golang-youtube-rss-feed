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
