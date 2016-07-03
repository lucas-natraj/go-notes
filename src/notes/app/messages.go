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