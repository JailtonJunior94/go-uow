// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Course struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type NewCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type NewCourse struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Category    *NewCategory `json:"category"`
}