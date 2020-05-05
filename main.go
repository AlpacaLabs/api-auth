package main

import (
	"sync"

	"github.com/AlpacaLabs/auth/internal/app"
	"github.com/AlpacaLabs/auth/internal/configuration"
)

func main() {
	c := configuration.LoadConfig()
	a := app.NewApp(c)

	var wg sync.WaitGroup

	wg.Add(1)
	go a.Run()

	wg.Wait()
}
