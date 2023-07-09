package main

import "fmt"

// the string attached to body will be json
type Note struct {
	Body  string `json:"body"`
}

func (a * App) SaveNote (n Note) string {
	return fmt.Sprintf("Saved Note (Note: %s)!", n.Body)
}