package main

const (
	WIDTH     = 800
	HEIGHT    = 800
	BASE_SIZE = 64
)

func main() {
	g := LoadGame("./assets/map-experiment.png")
	g.Run()
}
