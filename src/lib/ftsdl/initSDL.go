package ftsdl

import (
	"errors"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func initSSLErrorOverloading(err error) error {
	finalMsg := "SDL Initiation error : " + err.Error()
	return errors.New(finalMsg)
}

func InitSDL() (*sdl.Window, error) {
	// INIT SDL THE MODULE
	errInit := sdl.Init(sdl.INIT_EVERYTHING)
	if errInit != nil {
		return (nil), (errInit)
	}
	// CREATE THE WINDOW
	window, errCreateWindow := sdl.CreateWindow(
		"main",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		1264,
		712,
		sdl.WINDOW_SHOWN,
	)
	if errCreateWindow != nil {
		return (nil), (initSSLErrorOverloading(errCreateWindow))
	}
	// CREATE A RENDERER -> LIKE AN EMPTY FRAME
	renderer, errCreateRend := sdl.CreateRenderer(
		window,
		-1,
		sdl.RENDERER_ACCELERATED,
	)
	if errCreateRend != nil {
		return (nil), (initSSLErrorOverloading(errCreateRend))
	}
	defer renderer.Destroy()
	// LOAD PNG IMAGE
	image, errLoadImg := img.Load("./assets/bg.png")
	if errLoadImg != nil {
		return (nil), (initSSLErrorOverloading(errLoadImg))
	}
	texture, errCreateTex := renderer.CreateTextureFromSurface(image)
	if errCreateTex != nil {
		return (nil), (initSSLErrorOverloading(errCreateTex))
	}
	// INIT TTF RENDER MODULE
	errTtf := ttf.Init()
	if errTtf != nil {
		return (nil), (initSSLErrorOverloading(errTtf))
	}
	defer ttf.Quit()
	// LOAD FONT
	normalFont, errOpenFont := ttf.OpenFont("./assets/retroGaming.ttf", 20)
	if errOpenFont != nil {
		return (nil), (initSSLErrorOverloading(errOpenFont))
	}
	defer normalFont.Close()
	// LOAD FONT
	bigFont, errOpenBFont := ttf.OpenFont("./assets/retroGaming.ttf", 50)
	if errOpenBFont != nil {
		return (nil), (initSSLErrorOverloading(errOpenBFont))
	}
	defer bigFont.Close()
	// SET TEXT COLOR
	exitTextColor := sdl.Color{R: 0xC8, G: 0x00, B: 0x08, A: 255}
	// CLEAR THE RENDERER
	renderer.Clear()
	// COPY THE TEXTURE TO THE RENDERER
	renderer.Copy(texture, nil, nil)
	// RENDER TEXTS
	RenderText(
		renderer,
		"Close the window or press Esc to end the training.",
		normalFont,
		exitTextColor,
		20,
		20,
	)
	welcomeTextColor := sdl.Color{R: 0xFF, G: 0xFF, B: 0xFF, A: 235}
	RenderText(
		renderer,
		"Welcome in FT_ALITY !",
		bigFont,
		welcomeTextColor,
		275,
		588,
	)
	// ACTUALLY DRAW THE WINDOW
	renderer.Present()
	return (window), (nil)
}
