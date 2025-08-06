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

func cors(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
    w.Header().Set("Vary", "Origin")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

    if r.Method == http.MethodOptions {
      w.WriteHeader(http.StatusNoContent)
      return
    }
    next.ServeHTTP(w, r)
  })
}

func main() {
	mux := http.NewServeMux()
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

	mux.HandleFunc("/channel", handleChannelProfile)
	mux.HandleFunc("/video", handleVideoProfile)
	mux.HandleFunc("/category", handleCategoryProfile)
	mux.HandleFunc("/channelCategory", handleChannelCategoryProfile)
	mux.HandleFunc("/videoCategory", handleVideoCategoryProfile)
	mux.HandleFunc("/feed", handleFeedProfile)
	mux.HandleFunc("/channelFeed", handleChannelFeedProfile)
	mux.HandleFunc("/channelVideos", handleChannelVideosProfile)

	log.Println("Server is on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", cors(mux)))
}
