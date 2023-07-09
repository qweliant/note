package app

import (
	"context"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods 
// on startup load data from protobuff and begin I/O
func (a * App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// always open stream to save before close
func (a * App) Shutdown(ctx context.Context) {
}
