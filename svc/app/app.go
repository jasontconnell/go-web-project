package app

import "github.com/jasontconnell/go-web-project/data"

type App struct {
	repo    data.Repos
	devmode bool
}

func NewApp(r data.Repos, devmode bool) *App {
	a := &App{
		repo:    r,
		devmode: devmode,
	}
	return a
}
