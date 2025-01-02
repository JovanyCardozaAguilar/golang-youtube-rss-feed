package main

import (
	"context"
	"docker-go-youtube-feed/data"
	"docker-go-youtube-feed/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func handleClientProfile(w http.ResponseWriter, r *http.Request) {
	var clientId = r.URL.Query().Get("clientId")
	id, _ := strconv.Atoi(clientId)
	clientProfile, ok := data.GetUser(pool, ctx, id)
	if ok != nil || clientId == "" {
		http.Error(w, "ClientID does not exist Forbidden", http.StatusForbidden)
		return
	}
	r = r.WithContext(context.WithValue(r.Context(), "clientProfile", clientProfile))
	switch r.Method {
	case http.MethodGet:
		GetClientProfile(w, r)
	case http.MethodPatch:
		UpdateClientProfile(w, r)
	case http.MethodPut:
		PutClientProfile(w, r)
	case http.MethodDelete:
		DeleteClientProfile(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetClientProfile(w http.ResponseWriter, r *http.Request) {
	clientProfile := r.Context().Value("clientProfile").(*models.ClientProfile)

	w.Header().Set("Content-Type", "application/json")

	response := models.ClientProfile{
		Id:        clientProfile.Id,
		FirstName: clientProfile.FirstName,
		LastName:  clientProfile.LastName,
		Token:     clientProfile.Token,
	}
	json.NewEncoder(w).Encode(response)
}

func UpdateClientProfile(w http.ResponseWriter, r *http.Request) {
	clientProfile := r.Context().Value("clientProfile").(*models.ClientProfile)

	// Decode the JSON payload into struct
	var payloadData models.ClientProfile
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	clientProfile.FirstName = payloadData.FirstName
	clientProfile.LastName = payloadData.LastName
	clientProfile.Token = payloadData.Token
	fmt.Println("The payload data: ", payloadData)
	fmt.Println("The changed Client Profile: ", clientProfile)
	data.UpdateUser(pool, ctx, clientProfile.Id, *clientProfile)

	w.WriteHeader(http.StatusOK)
}

func PutClientProfile(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON payload into struct
	var payloadData models.ClientProfile
	if err := json.NewDecoder(r.Body).Decode(&payloadData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Println("The payload data: ", payloadData)
	data.InsertUser(pool, ctx, payloadData)

	w.WriteHeader(http.StatusOK)
}

func DeleteClientProfile(w http.ResponseWriter, r *http.Request) {
	clientProfile := r.Context().Value("clientProfile").(*models.ClientProfile)
	data.DeleteUser(pool, ctx, clientProfile.Id)

	w.WriteHeader(http.StatusOK)
}
