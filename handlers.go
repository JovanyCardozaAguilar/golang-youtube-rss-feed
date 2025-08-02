package main

import (
	"context"
	"docker-go-youtube-feed/data"
	"docker-go-youtube-feed/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handleChannelProfile(w http.ResponseWriter, r *http.Request) {
	if (r.Method == http.MethodPut) {
		PutChannelProfile(w, r)
		return
	}

	channelId := r.URL.Query().Get("channelId")
	channelProfile, ok := data.GetChannel(pool, ctx, channelId)
	r = r.WithContext(context.WithValue(r.Context(), "channelProfile", channelProfile))
	if ok != nil || channelId == "" {
		http.Error(w, "ChannelID does not exist Forbidden", http.StatusForbidden)
		return
	}

	switch r.Method {
	case http.MethodGet:
		GetChannelProfile(w, r)
	case http.MethodPatch:
		UpdateChannelProfile(w, r)
	case http.MethodDelete:
		DeleteChannelProfile(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetChannelProfile(w http.ResponseWriter, r *http.Request) {
	channelProfile := r.Context().Value("channelProfile").(*models.ChannelProfile)

	w.Header().Set("Content-Type", "application/json")

	response := models.ChannelProfile{
		ChannelId:	channelProfile.ChannelId,
		Username:	channelProfile.Username,
		Avatar:	channelProfile.Avatar,
	}
	json.NewEncoder(w).Encode(response)
}

func UpdateChannelProfile(w http.ResponseWriter, r *http.Request) {
	channelProfile := r.Context().Value("channelProfile").(*models.ChannelProfile)

	// Decode the JSON payload into struct
	var payloadData models.ChannelProfile
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		log.Println("JSON Decode Error:", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	channelProfile.Username = payloadData.Username
	channelProfile.Avatar = payloadData.Avatar
	fmt.Println("The payload data: ", payloadData)
	fmt.Println("The changed Channel Profile: ", channelProfile)
	err := data.UpdateChannel(pool, ctx, channelProfile.ChannelId, *channelProfile)

	if (err != nil) { 
		http.Error(w, "Channel Update error", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func PutChannelProfile(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON payload into struct
	var payloadData models.ChannelProfile
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		log.Println("JSON Decode Error:", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Println("The payload data: ", payloadData)
	err := data.InsertChannel(pool, ctx, payloadData)

	if (err != nil) { 
		http.Error(w, "Channel Put error", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteChannelProfile(w http.ResponseWriter, r *http.Request) {
	channelProfile := r.Context().Value("channelProfile").(*models.ChannelProfile)
	err := data.DeleteChannel(pool, ctx, channelProfile.ChannelId)

	if (err != nil) { 
		http.Error(w, "Channel Delete error", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleVideoProfile(w http.ResponseWriter, r *http.Request) {
	if (r.Method == http.MethodPut) {
		PutVideoProfile(w, r)
		return
	}

	videoId := r.URL.Query().Get("videoId")
	videoProfile, ok := data.GetVideo(pool, ctx, videoId)
	r = r.WithContext(context.WithValue(r.Context(), "videoProfile", videoProfile))

	if ok != nil || videoId == "" {
		http.Error(w, "VideoID does not exist Forbidden", http.StatusForbidden)
		return
	}

	switch r.Method {
	case http.MethodGet:
		GetVideoProfile(w, r)
	case http.MethodPatch:
		UpdateVideoProfile(w, r)
	case http.MethodDelete:
		DeleteVideoProfile(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetVideoProfile(w http.ResponseWriter, r *http.Request) {
	videoProfile := r.Context().Value("videoProfile").(*models.VideoProfile)

	w.Header().Set("Content-Type", "application/json")

	response := models.VideoProfile{
		VideoId:	videoProfile.VideoId,
		VChannelId:	videoProfile.VChannelId,
		Title:	videoProfile.Title,
		Thumbnail:	videoProfile.Thumbnail,
		Watched:	videoProfile.Watched,
	}
	json.NewEncoder(w).Encode(response)
}

func UpdateVideoProfile(w http.ResponseWriter, r *http.Request) {
	videoProfile := r.Context().Value("videoProfile").(*models.VideoProfile)

	// Decode the JSON payload into struct
	var payloadData models.VideoProfile
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		log.Println("JSON Decode Error:", err)
		log.Println("JSON Decode Error:", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	videoProfile.Title = payloadData.Title
	videoProfile.Thumbnail = payloadData.Thumbnail
	videoProfile.Watched = payloadData.Watched
	fmt.Println("The payload data: ", payloadData)
	fmt.Println("The changed Video Profile: ", videoProfile)
	err := data.UpdateVideo(pool, ctx, videoProfile.VideoId, *videoProfile)

	if (err != nil) { 
		http.Error(w, "Video Update error", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func PutVideoProfile(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON payload into struct
	var payloadData models.VideoProfile
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		log.Println("JSON Decode Error:", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Println("The payload data: ", payloadData)
	err := data.InsertVideo(pool, ctx, payloadData)

	if (err != nil) { 
		http.Error(w, "Video Put error", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteVideoProfile(w http.ResponseWriter, r *http.Request) {
	videoProfile := r.Context().Value("videoProfile").(*models.VideoProfile)
	err := data.DeleteVideo(pool, ctx, videoProfile.VideoId)

	if (err != nil) { 
		http.Error(w, "Channel Delete error", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleCategoryProfile(w http.ResponseWriter, r *http.Request) {
	if (r.Method == http.MethodPut) {
		PutCategoryProfile(w, r)
		return
	}

	categoryId := r.URL.Query().Get("categoryId")
	categoryProfile, ok := data.GetCategory(pool, ctx, categoryId)
	r = r.WithContext(context.WithValue(r.Context(), "categoryProfile", categoryProfile))

	if ok != nil || categoryId == "" {
		http.Error(w, "CategoryID does not exist Forbidden", http.StatusForbidden)
		return
	}
	
	switch r.Method {
	case http.MethodGet:
		GetCategoryProfile(w, r)
	case http.MethodPatch:
		UpdateCategoryProfile(w, r)
	case http.MethodDelete:
		DeleteCategoryProfile(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetCategoryProfile(w http.ResponseWriter, r *http.Request) {
	categoryProfile := r.Context().Value("categoryProfile").(*models.CategoryProfile)

	w.Header().Set("Content-Type", "application/json")

	response := models.CategoryProfile{
		CategoryId:	categoryProfile.CategoryId,
		CatName:	categoryProfile.CatName,
	}
	json.NewEncoder(w).Encode(response)
}

func UpdateCategoryProfile(w http.ResponseWriter, r *http.Request) {
	categoryProfile := r.Context().Value("categoryProfile").(*models.CategoryProfile)

	// Decode the JSON payload into struct
	var payloadData models.CategoryProfile
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		log.Println("JSON Decode Error:", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	categoryProfile.CatName = payloadData.CatName 
	fmt.Println("The payload data: ", payloadData)
	fmt.Println("The changed Category Profile: ", categoryProfile)
	err := data.UpdateCategory(pool, ctx, categoryProfile.CategoryId, *categoryProfile)

	if (err != nil) { 
		http.Error(w, "Channel Update error", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func PutCategoryProfile(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON payload into struct
	var payloadData models.CategoryProfile
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		log.Println("JSON Decode Error:", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Println("The payload data: ", payloadData)
	err := data.InsertCategory(pool, ctx, payloadData)

	if (err != nil) { 
		http.Error(w, "Category Put error", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteCategoryProfile(w http.ResponseWriter, r *http.Request) {
	categoryProfile := r.Context().Value("categoryProfile").(*models.CategoryProfile)
	err := data.DeleteCategory(pool, ctx, categoryProfile.CategoryId)

	if (err != nil) { 
		http.Error(w, "Channel Delete error", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleChannelCategoryProfile(w http.ResponseWriter, r *http.Request) {
	if (r.Method == http.MethodPut) {
		PutChannelCategoryProfile(w, r)
		return
	}
	if (r.Method == http.MethodDelete) {
		DeleteChannelCategoryProfile(w, r)
		return
	}

	channelCategoryId := r.URL.Query().Get("channelCategoryId")
	channelCategoryProfile, ok := data.GetChannelCategory(pool, ctx, channelCategoryId)
	r = r.WithContext(context.WithValue(r.Context(), "channelCategoryProfile", channelCategoryProfile))
	if ok != nil || channelCategoryId == "" {
		http.Error(w, "channelCategoryID does not exist Forbidden", http.StatusForbidden)
		return
	}
	if len(channelCategoryProfile) == 0 {
		http.Error(w, "No channel-category found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		GetChannelCategoryProfile(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetChannelCategoryProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(r.Context().Value("channelCategoryProfile"))
}

func PutChannelCategoryProfile(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON payload into struct
	var payloadData models.ChannelCategoryProfile
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		log.Println("JSON Decode Error:", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Println("The payload data: ", payloadData)
	err := data.InsertChannelCategory(pool, ctx, payloadData)

	if (err != nil) { 
		http.Error(w, "channel-category Put error", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteChannelCategoryProfile(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON payload into struct
	var payloadData models.ChannelCategoryProfile
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		log.Println("JSON Decode Error:", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Validate required field in JSON
	if payloadData.CcChannelId == "" || payloadData.CcCategoryId == "" {
		http.Error(w, "Missing ccChannelId or ccCategoryId in body", http.StatusBadRequest)
		return
	}

	fmt.Println("The payload data: ", payloadData)
	err := data.DeleteChannelCategory(pool, ctx, payloadData)

	if (err != nil) { 
		http.Error(w, "channel-category delete error", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleVideoCategoryProfile(w http.ResponseWriter, r *http.Request) {
	if (r.Method == http.MethodPut) {
		PutVideoCategoryProfile(w, r)
		return
	}
	if (r.Method == http.MethodDelete) {
		DeleteVideoCategoryProfile(w, r)
		return
	}

	videoCategoryId := r.URL.Query().Get("videoCategoryId")
	videoCategoryProfile, ok := data.GetVideoCategory(pool, ctx, videoCategoryId)
	r = r.WithContext(context.WithValue(r.Context(), "videoCategoryProfile", videoCategoryProfile))
	if ok != nil || videoCategoryId == "" {
		http.Error(w, "videoCategoryID does not exist Forbidden", http.StatusForbidden)
		return
	}
	if len(videoCategoryProfile) == 0 {
		http.Error(w, "No video-category found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		GetVideoCategoryProfile(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetVideoCategoryProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(r.Context().Value("videoCategoryProfile"))
}

func PutVideoCategoryProfile(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON payload into struct
	var payloadData models.VideoCategoryProfile
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		log.Println("JSON Decode Error:", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Println("The payload data: ", payloadData)
	err := data.InsertVideoCategory(pool, ctx, payloadData)

	if (err != nil) { 
		http.Error(w, "video-category Put error", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteVideoCategoryProfile(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON payload into struct
	var payloadData models.VideoCategoryProfile
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		log.Println("JSON Decode Error:", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Validate required field in JSON
	if payloadData.VcVideoId == "" || payloadData.VcCategoryId == "" {
		http.Error(w, "Missing vcVideoId or vcCategoryId in body", http.StatusBadRequest)
		return
	}

	fmt.Println("The payload data: ", payloadData)
	err := data.DeleteVideoCategory(pool, ctx, payloadData)

	if (err != nil) { 
		http.Error(w, "video-category delete error", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}


func handleFeedProfile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetFeedProfile(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetFeedProfile(w http.ResponseWriter, r *http.Request) {
	feedProfile, ok := data.GetFeed(pool, ctx)
	r = r.WithContext(context.WithValue(r.Context(), "feedProfile", feedProfile))
	if ok != nil {
		http.Error(w, "Error with getting feed", http.StatusForbidden)
		return
	}
	if len(feedProfile) == 0 {
		http.Error(w, "No videos found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(r.Context().Value("feedProfile"))
}

func handleYouTubeProfile(w http.ResponseWriter, r *http.Request) {
	youtubeId := r.URL.Query().Get("youtubeId")
	youtubeProfile, ok := data.GetVideoCategory(pool, ctx, youtubeId)
	r = r.WithContext(context.WithValue(r.Context(), "videoCategoryProfile", youtubeProfile))
	if ok != nil || youtubeId == "" {
		http.Error(w, "videoCategoryID does not exist Forbidden", http.StatusForbidden)
		return
	}
	if len(youtubeProfile) == 0 {
		http.Error(w, "No video-category found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		GetVideoCategoryProfile(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetYouTubeProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(r.Context().Value("videoCategoryProfile"))
}

