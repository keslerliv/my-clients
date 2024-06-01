package handlers

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/keslerliv/my-clients/internal/entities"
	"github.com/keslerliv/my-clients/internal/models"
)

// Client GET handler
func ClientGet(w http.ResponseWriter, r *http.Request) {

	// get id from request
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Printf("error parsing id: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get client by id model
	client, err := models.ClientGet(int64(id))
	if err != nil {
		log.Printf("error getting client: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// http response
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(client)
}

// Client GET-ALL handler
func ClientList(w http.ResponseWriter, r *http.Request) {

	// get all clients model
	clients, err := models.ClientGetAll()
	if err != nil {
		log.Printf("error getting clients: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// http response
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clients)
}

// Client POST handler
func ClientCreate(w http.ResponseWriter, r *http.Request) {

	// decode client request body
	var client entities.Client
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		log.Printf("error decoding client: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// insert client model
	id, err := models.ClientInsert(client)
	var response map[string]any
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response = map[string]any{"id": id}

	// http response
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Client PUT handler
func ClientUpdate(w http.ResponseWriter, r *http.Request) {

	// get id from request
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Printf("error parsing id: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// decode client request body
	var client entities.Client
	err = json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		log.Printf("error decoding client: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// update client model
	_, err = models.ClientUpdate(int64(id), client)
	if err != nil {
		log.Printf("error updating client: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]any{"id": id}

	// http response
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Client DELETE handler
func ClientDelete(w http.ResponseWriter, r *http.Request) {

	// get id from request
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Printf("error parsing id: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// delete client model
	_, err = models.ClientDelete(int64(id))
	if err != nil {
		log.Printf("error deleting client: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]any{"id": id}

	// http response
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Create clients from TXT file
func CreateClientsFromTXT(w http.ResponseWriter, r *http.Request) {

	// parse request
	err := r.ParseMultipartForm(20 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get file from request
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// read file and get clients
	scanner := bufio.NewScanner(file)
	var clients []entities.Client

	// skip header line
	scanner.Scan()

	// map lines and append to clients
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		// parse date field
		var dateLastPurchase *time.Time
		if fields[3] != "NULL" {
			date, err := time.Parse("2006-01-02", fields[3])
			if err == nil {
				dateLastPurchase = &date
			}
		}

		// parse int fields
		AverageTicket, _ := strconv.ParseInt(strings.ReplaceAll(fields[4], ",", ""), 10, 64)
		TicketLastPurchase, _ := strconv.ParseInt(strings.ReplaceAll(fields[5], ",", ""), 10, 64)

		// parse string fields
		var frequentStore *string
		if fields[6] != "NULL" {
			frequentStore = &fields[6]
		}
		var lastStore *string
		if fields[7] != "NULL" {
			lastStore = &fields[7]
		}

		// add client
		client := entities.Client{
			CPF:                fields[0],
			Private:            fields[1] == "1",
			Incomplete:         fields[2] == "1",
			DateLastPurchase:   dateLastPurchase,
			AverageTicket:      AverageTicket,
			TicketLastPurchase: TicketLastPurchase,
			FrequentStore:      frequentStore,
			LastStore:          lastStore,
		}
		clients = append(clients, client)
	}

	// insert client model
	err = models.ClientListInsert(clients)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// http response
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("append clients from txt file successfully")
}
