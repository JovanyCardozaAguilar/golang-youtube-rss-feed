package main

import (
	"context"
	"docker-go-youtube-feed/data"
	"docker-go-youtube-feed/models"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	ctx  context.Context
	pool *models.Postgres
)

func main() {
	dsn := "postgres://testUser1:password@localhost:5432/testdb1?sslmode=disable"
	ctx = context.Background()
	var err error
	pool, err = data.CreateDBPool(ctx, dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("The Pool: ", pool)
	fmt.Println(data.QueryGreeting(ctx, pool))
	// Test channels
	fmt.Println(data.QuerySingleTestChannel(ctx, pool))
	accounts, _ := data.QueryMultiTestChannel(ctx, pool)
	for _, account := range accounts {
		fmt.Printf("%#v\n", account)
	}
	// Test videos
	fmt.Println(data.QuerySingleTestVideo(ctx, pool))
	videos, _ := data.QueryMultiTestVideo(ctx, pool)
	for _, video := range videos {
		fmt.Printf("%#v\n", video)
	}

	// Test videos
	fmt.Println(data.QuerySingleTestCategory(ctx, pool))
	categories, _ := data.QueryMultiTestCategory(ctx, pool)
	for _, category := range categories {
		fmt.Printf("%#v\n", category)
	}

	http.HandleFunc("/channel", handleChannelProfile)

	log.Println("Server is on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
