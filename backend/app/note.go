package app

import (
	"fmt"
)

func (a *App) SaveNote(n *Note) string {
	return fmt.Sprintf("Saved Note (Note: %s)!", n.Body)
}

func (a *App) ReadNote(n *Note) *Note {
	return n
}
