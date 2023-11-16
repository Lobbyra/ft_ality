package ftsdl

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func RenderText(
	renderer *sdl.Renderer,
	text string,
	font *ttf.Font,
	color sdl.Color,
	x int32,
	y int32,
) {
	// CREATE A SURFACE WITH THE TEXT DRAWN IN IT
	surface, err := font.RenderUTF8Solid(text, color)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to render text: %s\n", err)
		return
	}
	defer surface.Free()
	// TRANSFORM THE SURFACE IN A TEXTURE
	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		return
	}
	defer texture.Destroy()
	// DRAW THE TEXTURE IN THE RENDER
	rect := &sdl.Rect{X: x, Y: y, W: surface.W, H: surface.H}
	renderer.Copy(texture, nil, rect)
}
