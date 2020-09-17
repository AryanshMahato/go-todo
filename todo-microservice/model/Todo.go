package model

// Todo
// This struct will work as Model of Todo in database
type Todo struct {
	Title       string `json:"title"`
	IsCompleted bool   `json:"is_completed"`
}
