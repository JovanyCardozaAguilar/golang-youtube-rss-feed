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
	var channelId = r.URL.Query().Get("channelId")
	channelProfile, ok := data.GetChannel(pool, ctx, channelId)
	if ok != nil || channelId == "" {
		http.Error(w, "ChannelID does not exist Forbidden", http.StatusForbidden)
		return
	}
	r = r.WithContext(context.WithValue(r.Context(), "channelProfile", channelProfile))
	switch r.Method {
	case http.MethodGet:
		GetChannelProfile(w, r)
	case http.MethodPatch:
		UpdateChannelProfile(w, r)
	case http.MethodPut:
		PutChannelProfile(w, r)
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
	data.UpdateChannel(pool, ctx, channelProfile.ChannelId, *channelProfile)

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
	data.InsertChannel(pool, ctx, payloadData)

	w.WriteHeader(http.StatusOK)
}

func DeleteChannelProfile(w http.ResponseWriter, r *http.Request) {
	channelProfile := r.Context().Value("channelProfile").(*models.ChannelProfile)
	data.DeleteChannel(pool, ctx, channelProfile.ChannelId)

	w.WriteHeader(http.StatusOK)
}

func handleVideoProfile(w http.ResponseWriter, r *http.Request) {
	var videoId = r.URL.Query().Get("videoId")
	videoProfile, ok := data.GetVideo(pool, ctx, videoId)
	if ok != nil || videoId == "" {
		http.Error(w, "VideoID does not exist Forbidden", http.StatusForbidden)
		return
	}
	r = r.WithContext(context.WithValue(r.Context(), "videoProfile", videoProfile))
	switch r.Method {
	case http.MethodGet:
		GetVideoProfile(w, r)
	case http.MethodPatch:
		UpdateVideoProfile(w, r)
	case http.MethodPut:
		PutVideoProfile(w, r)
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
		Title:	videoProfile.Title,
		Thumbnail:	videoProfile.Thumbnail,
		Watched:	videoProfile.Watched,
		VideoChannel:	videoProfile.VideoChannel,
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
	data.UpdateVideo(pool, ctx, videoProfile.VideoId, *videoProfile)

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
	data.InsertVideo(pool, ctx, payloadData)

	w.WriteHeader(http.StatusOK)
}

func DeleteVideoProfile(w http.ResponseWriter, r *http.Request) {
	videoProfile := r.Context().Value("videoProfile").(*models.VideoProfile)
	data.DeleteVideo(pool, ctx, videoProfile.VideoId)

	w.WriteHeader(http.StatusOK)
}

func handleCategoryProfile(w http.ResponseWriter, r *http.Request) {
	var categoryId = r.URL.Query().Get("categoryId")
	categoryProfile, ok := data.GetCategory(pool, ctx, categoryId)
	if ok != nil || categoryId == "" {
		http.Error(w, "CategoryID does not exist Forbidden", http.StatusForbidden)
		return
	}
	r = r.WithContext(context.WithValue(r.Context(), "categoryProfile", categoryProfile))
	switch r.Method {
	case http.MethodGet:
		GetCategoryProfile(w, r)
	case http.MethodPatch:
		UpdateCategoryProfile(w, r)
	case http.MethodPut:
		PutCategoryProfile(w, r)
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
		CatChannel:	categoryProfile.CatChannel,
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

	categoryProfile.CategoryId = payloadData.CategoryId 
	categoryProfile.CatName = payloadData.CatName 
	categoryProfile.CatChannel = payloadData.CatChannel 
	fmt.Println("The payload data: ", payloadData)
	fmt.Println("The changed Category Profile: ", categoryProfile)
	data.UpdateCategory(pool, ctx, categoryProfile.CategoryId, *categoryProfile)

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
	data.InsertCategory(pool, ctx, payloadData)

	w.WriteHeader(http.StatusOK)
}

func DeleteCategoryProfile(w http.ResponseWriter, r *http.Request) {
	categoryProfile := r.Context().Value("categoryProfile").(*models.CategoryProfile)
	data.DeleteCategory(pool, ctx, categoryProfile.CategoryId)

	w.WriteHeader(http.StatusOK)
}

