package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/Rulox/ebitmx"
	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	// Our map is 10x10 tiles, with 16 px per tile
	screenWidth  = 160
	screenHeight = 160
)

type Game struct {
	myMap *ebitmx.EbitenMap
	tiles *ebiten.Image
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw map using the same example as
	// https://ebiten.org/examples/tiles.html

	// Get the Tiles image width and height and divide by the Tile size to get the number of
	// tiles in width and height
	tileWidth, tileHeight := g.tiles.Size()
	tileWidth = tileWidth / g.myMap.TileWidth
	tileHeight = tileHeight / g.myMap.TileHeight

	for _, l := range g.myMap.Layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}

			op.GeoM.Translate(float64((i%g.myMap.MapWidth)*g.myMap.TileWidth), float64((i/g.myMap.MapHeight)*g.myMap.TileHeight))
			// Transform from 1D slice to 2 coordinates
			sx := ((t % tileWidth) - 1) * g.myMap.TileWidth
			sy := (t / tileWidth) * g.myMap.TileHeight

			// Draw the tile in the position "t"
			screen.DrawImage(g.tiles.SubImage(image.Rect(sx, sy, sx+g.myMap.TileHeight, sy+g.myMap.TileWidth)).(*ebiten.Image), op)
		}
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Game (Demo)")

	tiles, _, err := ebitenutil.NewImageFromFile("overworld.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	myMap, err := ebitmx.GetEbitenMap("map.tmx")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	game := &Game{
		tiles: tiles,
		myMap: myMap,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
