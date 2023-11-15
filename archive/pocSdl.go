package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	rect := sdl.Rect{0, 0, 200, 200}
	colour := sdl.Color{R: 255, G: 0, B: 255, A: 255} // purple
	pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	surface.FillRect(&rect, pixel)
	window.UpdateSurface()

	joystick := sdl.JoystickOpen(0)
	if joystick == nil {
		fmt.Println("Impossible d'ouvrir la manette:", sdl.GetError())
		return
	}
	defer joystick.Close()
	println(sdl.NumJoysticks())

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				joystick.Close()
				running = false
				break
			case *sdl.JoyAxisEvent:
				// Gestion des événements d'axe de la manette
				axisEvent := event.(*sdl.JoyAxisEvent)
				if axisEvent.Axis == 0 || (axisEvent.Value > -10000 && axisEvent.Value < 10000) {
					break
				}
				surface.FillRect(nil, 0)
				if axisEvent.Axis == 2 { // X
					rect.X = rect.X + (int32(axisEvent.Value) / 10000)
				} else if axisEvent.Axis == 3 { //X
					rect.Y = rect.Y + (int32(axisEvent.Value) / 10000)
				}
				surface.FillRect(&rect, pixel)
				window.UpdateSurface()

				// fmt.Printf("Axe %d: %d (%d) [x: %d, y: %d] {%d}\n", axisEvent.Axis, axisEvent.Value, axisEvent.Timestamp, rect.X, rect.Y, (int32(axisEvent.Value) / 10000))
			case *sdl.JoyButtonEvent:
				buttonEvent := event.(*sdl.JoyButtonEvent)
				fmt.Printf("Bouton %d: %d\n", buttonEvent.Button, buttonEvent.State)
			default:
				break
			}
		}
	}
}
