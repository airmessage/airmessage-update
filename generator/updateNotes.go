package main

// UpdateNotes represents update notes for a specific language
type UpdateNotes struct {
	Lang string `json:"lang"`
	Message string `json:"message"`
}
