package app

import (
	"fmt"
)

type App struct {
	Id string
}

func (a *App) Identify() {
	fmt.Println(a.Id)
}
