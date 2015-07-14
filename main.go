package main

import "github.com/veandco/go-sdl2/sdl"

const (
	windowWidth  = 800
	windowHeight = 600
)

func main() {
	sdl.Init(sdl.INIT_VIDEO)
	defer sdl.Quit()
	window, err := sdl.CreateWindow("Tank", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, windowWidth, windowHeight, 0)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

mainloop:
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				break mainloop
			}
		}
		windowSurface, _ := window.GetSurface()
		windowSurface.FillRect(nil, sdl.MapRGB(windowSurface.Format, 0, 0, 0))

		window.UpdateSurface()
		sdl.Delay(1000)
	}
}
