package main

import "github.com/nikitaSstepanov/templates/golang/internal/app"

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	a := app.New()

	a.Run()
}
