package main

import (
	"context"
	"fmt"
)

type CommonResponse struct {
	Code   int
	ErrMsg string
	Value  interface{}
}

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
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Login(name, password string) CommonResponse {
	var ret CommonResponse
	fmt.Printf("cal Login, name=%s, password=%s\n", name, password)
	return ret
}
