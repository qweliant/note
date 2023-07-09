package app

import (
	"fmt"
)

// the string attached to body will be json
type EditorNote struct {
	Body string `json:"body"`
}

func (a *App) SaveNote(n EditorNote) string {
	return fmt.Sprintf("Saved Note (Note: %s)!", n.Body)
}

func (a *App) ReadNote(n EditorNote) EditorNote {
	return n
}
