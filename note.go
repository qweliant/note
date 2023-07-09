package main

import "fmt"

// the string attached to body will be json
type Note struct {
	Body  string `json:"body"`
}

func (a *App) CreateNote(n Note) string {
	return fmt.Sprintf("Contents Are (Age: %s)!", n.Body)

}
