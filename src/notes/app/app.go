// Notes Service.
//
// Notes management service
//
// Schemes: http
// Host: localhost:3000
// BasePath: /
// Version: 1.0.0
//
// Produces:
// - application/json
//
//
// swagger:meta
//
// go:generate swagger generate spec -o swagger.json
package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/pat"
	"github.com/rs/cors"
)

// NotesService is the main service that manages the Notes
type NotesService struct {
	Notes []Note
}

// NewNotesService creates a new instance of the NotesService
func NewNotesService() *NotesService {
	t := new(NotesService)
	t.Notes = make([]Note, 0)
	return t
}

// Singleton to store all the notes
var note_service = NewNotesService()

// swagger:route GET /info info info-status
//
// Get service information
//
// Get service information
//
// Security:
//   my_auth: email
//
// Responses:
// 	200: InfoResponse
func (ns *NotesService) Info(w http.ResponseWriter, r *http.Request) {

	info := &Info{
		Service: "Notes",
		Status:  "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

// swagger:route POST /notes notes notes-add
//
// Add a new note
//
// Adds a new note
//
// Security:
//   my_auth: email
//
// Responses:
// 	200: Success
func (ns *NotesService) AddNote(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body AddNoteRequestBody
	err := decoder.Decode(&body)
	if err != nil {
		panic(nil)
	}

	var Note = Note{
		Title:   body.Title,
		Content: body.Content,
	}

	ns.Notes = append(note_service.Notes, Note)
}

// swagger:route GET /notes notes notes-fetchAll
//
// Fetches all the notes
//
// Returns the entire list of notes
//
// Security:
//   my_auth: email
//
// Responses:
// 	200: FetchAllNotesResponse
func (ns *NotesService) FetchAllNotes(w http.ResponseWriter, r *http.Request) {

	body := &FetchAllNotesResponseBody{
		Notes: ns.Notes,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

// swagger:route GET /notes/{id} notes notes-fetchOne
//
// Fetch a single note
//
// Fetch a single note by index
//
// Security:
//   my_auth: email
//
// Responses:
// 	200: FetchNoteResponse
func (ns *NotesService) GetNote(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	id := params.Get(":id")
	if len(id) == 0 {
		http.Error(w, "missing note id", http.StatusBadRequest)
		return
	}

	index, err := strconv.Atoi(id)
	if err != nil || (index < 0 || index >= len(ns.Notes)) {
		http.Error(w, "invalid note id", http.StatusBadRequest)
		return
	}

	body := &FetchNoteResponse{
		Body: &Note{
			Title:   ns.Notes[index].Title,
			Content: ns.Notes[index].Content,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

// swagger:route DELETE /notes/{id} notes notes-remove
//
// Removes a single note
//
// Removes a single note by index
//
// Security:
//   my_auth: email
//
// Responses:
// 	200: Success
func (ns *NotesService) RemoveNote(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	id := params.Get(":id")
	if len(id) == 0 {
		http.Error(w, "missing note id", http.StatusBadRequest)
		return
	}

	index, err := strconv.Atoi(id)
	if err != nil || (index < 0 || index >= len(ns.Notes)) {
		http.Error(w, "invalid note id", http.StatusBadRequest)
		return
	}

	ns.Notes = append(ns.Notes[:index], ns.Notes[index+1:]...)
}

// swagger:route PUT /notes/{id} notes notes-update
//
// Updates a single note
//
// Updates the note at the specified index
//
// Security:
//   my_auth: email
//
// Responses:
// 	200: Success
func (ns *NotesService) UpdateNote(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	id := params.Get(":id")
	if len(id) == 0 {
		http.Error(w, "missing note id", http.StatusBadRequest)
		return
	}

	index, err := strconv.Atoi(id)
	if err != nil || (index < 0 || index >= len(ns.Notes)) {
		http.Error(w, "invalid note id", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var body AddNoteRequestBody
	err = decoder.Decode(&body)
	if err != nil {
		panic(nil)
	}

	ns.Notes[index] = Note{
		Title:   body.Title,
		Content: body.Content,
	}
}

// initialize the identity service endpoints
func init() {
	r := pat.New()
	r.Get("/info", note_service.Info)
	r.Get("/notes/{id}", note_service.GetNote)
	r.Delete("/notes/{id}", note_service.RemoveNote)
	r.Put("/notes/{id}", note_service.UpdateNote)
	r.Get("/notes", note_service.FetchAllNotes)
	r.Post("/notes", note_service.AddNote)

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
	})
	handler := c.Handler(r)
	http.Handle("/", handler)
}
