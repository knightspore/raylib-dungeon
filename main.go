package main

const (
	WIDTH     = 1920
	HEIGHT    = 1080
	BASE_SIZE = 64
)

func main() {
	g := LoadGame("./assets/map-experiment.png")
	g.Run()
}
