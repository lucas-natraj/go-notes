---
layout: post
category : documentation
author: Lucas Natraj
tags: [Swagger, YAML, JSON, Generate, Go, Golang]
title: Swagger Generation from a Go Web Application 
---

## Simple Go Web Application

Application Structure:
```
.  <---------------------------- GOPATH
├── bin
│   └── ...
├── pkg
│   └── ...
├── src
│   ├── ...
│   └── notes
│       ├── app
│       │   ├── app.go
│       │   └── notes.yaml
│       └── main.go
├── swagger-aux.json
└── swagger.json   <------------- Generated Swagger

```

The implementation below is a fairly straightforward note management application, with the following REST-ful api -
* POST /notes - Add a new note
* GET /notes - Get all notes
* GET /notes/:index - Get note at specified index
* PUT /notes/:index - Update note at specified index
* DELETE /notes/:index - Delete note at specified index

```go
// -------  src/notes/app/app.go  -------

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
	"net/http"
	"github.com/gorilla/pat"
	"encoding/json"
	"github.com/rs/cors"
	"strconv"
)

type NotesService struct {
	Notes []Note
}

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
// Responses:
// 	200: InfoResponse
func (ns *NotesService) Info(w http.ResponseWriter, r *http.Request) {

	info := &Info{
		Service: "Notes",
		Status: "ok",
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
		Title: body.Title,
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
			Title:ns.Notes[index].Title,
			Content:ns.Notes[index].Content,
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

	ns.Notes = append(ns.Notes[:index], ns.Notes[index + 1:]...)
}

// swagger:route PUT /notes/{id} notes notes-update
//
// Updates a single note
//
// Updates the note at the specified index
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
		Title: body.Title,
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
```

```go
// -------  src/notes/app/messages.go  -------

package app

type Note struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Info struct {
	Service string `json:"service"`
	Status  string `json:"status"`
}

// InfoResponse
//
// swagger:response
type InfoResponse struct {
	// required: true
	// in: body
	Body *Info `json:"body"`
}

// swagger:parameters notes-add
type AddNoteRequest struct {
	// Payload
	//
	// required: true
	// in: body
	Body *AddNoteRequestBody
}

// AddNoteRequestBody
// swagger:model
type AddNoteRequestBody struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// The response containing the list of Notes
//
// swagger:response
type FetchAllNotesResponse struct {
	// required: true
	// in: body
	Notes *FetchAllNotesResponseBody `json:"notes"`
}

type FetchAllNotesResponseBody struct {
	// The Note List
	Notes []Note `json:"notes"`
}

// swagger:parameters notes-fetchOne
type FetchNoteRequest struct {
	// Index of Note
	//
	// required: true
	// in: path
	// minimum: 0
	// default: 0
	NoteId int        `json:"id" binding:"required"`
}

// The response for a single note fetch
//
// swagger:response
type FetchNoteResponse struct {
	// required: true
	// in: body
	Body *Note `json:"body"`
}

// swagger:parameters notes-remove
type DeleteNoteRequest struct {
	// Index of Note
	//
	// required: true
	// in: path
	// minimum: 0
	// default: 0
	NoteId int        `json:"id" binding:"required"`
}

// swagger:parameters notes-update
type UpdateNoteRequest struct {
	// Index of Note
	//
	// required: true
	// in: path
	// minimum: 0
	// default: 0
	NoteId int        `json:"id" binding:"required"`

	// Payload
	//
	// required: true
	// in: body
	Body   *AddNoteRequestBody
}

// Success
//
// swagger:response
type Success struct {

}
```

Providing security definitions is done via an auxilliary file that is merged in during the spec generation process.  

`-------  swagger-aux.json  -------`
```json
{
  "securityDefinitions": {
    "my_auth": {
      "type": "oauth2",
      "flow": "implicit",
      "authorizationUrl": "https://accounts.google.com/o/oauth2/v2/auth",
      "scopes": {
        "email": "default scope"
      }
    }
  }
}
```

## Running Swagger UI Locally 
Instructions at https://github.com/swagger-api/swagger-ui

- To run the Swagger Editor locally using docker: 
```bash
docker pull swaggerapi/swagger-editor
docker run -d -p 80:8080 swaggerapi/swagger-editor 
```
Then, navigate to `http://locahost` (or your local docker ip)

- To run the Swagger Editor locally using Node:
```bash
$ git clone https://github.com/swagger-api/swagger-editor.git
$ cd swagger-editor
$ npm install
$ npm start
```

## Live-Reloading Go Web Applications

Use [gin](https://github.com/codegangsta/gin).

```bash
$ go get github.com/codegangsta/gin
$ $GOPATH/bin/gin
``` 

Note: The web application should query the PORT environment variable to know on which port to listen:
```go
// -------  src/notes/main.go  -------

package main

import (
	"net/http"
	"os"
)

func getenv(key string, fallback string) string {
	v := os.Getenv(key)
	if len(v) != 0 {
		return v
	}
	return fallback
}

// used for local testing / debugging
func main() {
	http.ListenAndServe(":" + getenv("PORT", "8080"), nil)
}
```

# Generating swagger documentation

Swagger generation fo Go applications uses [go-swagger](https://github.com/go-swagger/go-swagger).  
The documentation at https://goswagger.io describes the various go documentation tags that are required for generating swagger documentation.

The following commands should be executed at $GOPATH
```bash
$ go get -u github.com/go-swagger/go-swagger/cmd/swagger
$ ./bin/swagger generate spec --input=./swagger-aux.json --output=./swagger.json --base-path=./src/notes/app/
```

Note: The swagger editor does display some warnings for unused attributes and definitions, which I have not been able to eliminate. 
The good news is that these warnings have no noticeable impact on the resulting rendered swagger documentation page.  

### Links

* [App Engine with Go](https://cloud.google.com/appengine/docs/go/)
* [Uploading, Downloading, and Managing a Go App](https://cloud.google.com/appengine/docs/go/tools/uploadinganapp)
* [Go Swagger](https://github.com/go-swagger/go-swagger)
* [Go Swagger Toolkit / Documentation](https://goswagger.io)
* [Notes Application Source](https://github.com/lucas-natraj/go-notes)
