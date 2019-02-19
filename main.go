package main

import "budgetBookArch/app"

func main() {
	// Instantiates a new application, parses the CLI input and runs the
	// respective Command.
	app.New().Run()
}
