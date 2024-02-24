package main

func main() {
	app, cleanup, err := Initialize()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Serve(); err != nil {
		panic(err)
	}
}
