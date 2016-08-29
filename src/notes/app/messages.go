package app

// Note is the persistence format of a note
type Note struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Info is the response format for the Info request
type Info struct {
	Service string `json:"service"`
	Status  string `json:"status"`
}

// InfoResponse is the response to the Info request
//
// swagger:response
type InfoResponse struct {
	// required: true
	// in: body
	Body *Info `json:"body"`
}

// AddNoteRequest is the request structure
//
// swagger:parameters notes-add
type AddNoteRequest struct {
	// Payload
	//
	// required: true
	// in: body
	Body *AddNoteRequestBody
}

// AddNoteRequestBody is the body of the AddNoteRequest request
//
// swagger:model
type AddNoteRequestBody struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// FetchAllNotesResponse is the response containing the list of Notes
//
// swagger:response
type FetchAllNotesResponse struct {
	// required: true
	// in: body
	Notes *FetchAllNotesResponseBody `json:"notes"`
}

// FetchAllNotesResponseBody is the body of the FetchAllNotesResponse response
type FetchAllNotesResponseBody struct {
	// The Note List
	Notes []Note `json:"notes"`
}

// FetchNoteRequest is the request for the FetchNote request
// swagger:parameters notes-fetchOne
type FetchNoteRequest struct {
	// Index of Note
	//
	// required: true
	// in: path
	// minimum: 0
	// default: 0
	NoteID int `json:"id" binding:"required"`
}

// FetchNoteResponse is the response when fetching a specific note
//
// swagger:response
type FetchNoteResponse struct {
	// required: true
	// in: body
	Body *Note `json:"body"`
}

// DeleteNoteRequest is the request for deleting a note
//
// swagger:parameters notes-remove
type DeleteNoteRequest struct {
	// Index of the note to be removed
	//
	// required: true
	// in: path
	// minimum: 0
	// default: 0
	NoteID int `json:"id" binding:"required"`
}

// UpdateNoteRequest is the request for updating a specific note
//
// swagger:parameters notes-update
type UpdateNoteRequest struct {
	// Index of Note to be updated
	//
	// required: true
	// in: path
	// minimum: 0
	// default: 0
	NoteID int `json:"id" binding:"required"`

	// Payload
	//
	// required: true
	// in: body
	Body *AddNoteRequestBody
}

// Success is the response for a successful request
//
// swagger:response
type Success struct {
}
