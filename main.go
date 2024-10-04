package main

const (
	WIDTH     = 800
	HEIGHT    = 800
	BASE_SIZE = 64
)

func main() {
	g := LoadGameFromImage("./assets/map-experiment.png")
	g.Run()
}
