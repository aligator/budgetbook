package main

import "budgetBook/app"

func main() {
	// Instantiates a new application, and runs the web application.
	app.NewServer(":8080").Run()
}
