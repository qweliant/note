package main

import "fmt"

type Note struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (a *App) CreateNote(n Note) string {
	return fmt.Sprintf("Contents Are (Age: %s)!", n.Body)

}
