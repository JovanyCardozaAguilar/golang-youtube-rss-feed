package main

import (
	"context"
	"docker-go-youtube-feed/data"
	"docker-go-youtube-feed/models"
	"encoding/json"
	"fmt"
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
