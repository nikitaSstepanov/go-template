package main

import "app/internal/app"

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	a := app.New()

	a.Run()
}
