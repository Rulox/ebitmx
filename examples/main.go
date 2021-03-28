package main

import (
	"embed"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/Rulox/ebitmx"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// embeddedFS holds our game assets so we can distribute our game as a single binary
//go:embed map.tmx tileset.tsx overworld.png
var embeddedFS embed.FS

const (
	screenWidth  = 960
	screenHeight = 960
)

type Game struct {
	myMap           *ebitmx.EbitenMap
	tileset         *ebitmx.EbitenTileset
	atlas           *ebiten.Image
	currentTileType string
}

func (g *Game) Update() error {
	// Set the tile type to "off-screen" the cursor ever goes out of bounds
	cx, cy := ebiten.CursorPosition()
	if cx < 0 || cy < 0 || cx >= screenWidth || cy >= screenHeight {
		g.currentTileType = "off screen"
		return nil
	}

	// Find the id of the tile in the first layer that the cursor is currently over
	tx, ty := g.cursorPositionInTileSpace()
	id := g.myMap.Layers[0][(ty*g.myMap.MapWidth)+tx]

	// Unset tiles have an ID of 0 - handle them
	if id == 0 {
		g.currentTileType = "unset"
		return nil
	}

	// Only some of our tiles are named (and as a result, only some of them appear in our
	// .tsx file), so their position in g.tileset.Tiles[] isn't directly related to their id.
	// As such, we have to iterate over the list each time and find the tile we're looking
	// for. This could be improved by caching with a map of ids to *Tiles.
	found := false
	for _, t := range g.tileset.Tiles {
		if t.Id+1 == id {
			g.currentTileType = t.Type
			found = true
			break
		}
	}

	// We've come across a tile that asn't explicitly stored in the .tsx
	if !found {
		g.currentTileType = "implicit tile - not in .tsx file"
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw map using the same method as the official tiles example
	// https://ebiten.org/examples/tiles.html

	// The scaling we use is consistent across all tiles, so we'll
	// calculate it outside of the tile-drawing loop
	sx := float64(screenWidth / (g.myMap.MapWidth * g.tileset.TileWidth))
	sy := float64(screenHeight / (g.myMap.MapHeight * g.tileset.TileHeight))

	for _, l := range g.myMap.Layers {
		for i, id := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(
				float64((i%g.myMap.MapWidth)*g.myMap.TileWidth),
				float64((i/g.myMap.MapHeight)*g.myMap.TileHeight),
			)
			op.GeoM.Scale(sx, sy)

			screen.DrawImage(g.getTileImgByID(id), op)
		}
	}

	cx, cy := ebiten.CursorPosition()
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("cx:%d, cy:%d\ntype: %s\n", cx, cy, g.currentTileType),
	)
}

func (g *Game) Layout(ow, oh int) (int, int) {
	return ow, oh
}

func (g *Game) getTileImgByID(id int) *ebiten.Image {
	// The tsx format starts counting tiles from 1, so to make these calculations
	// work correctly, we need to decrement the ID by 1
	id -= 1

	x0 := (id % g.tileset.TilesetWidth) * g.tileset.TileWidth
	y0 := (id / g.tileset.TilesetWidth) * g.tileset.TileHeight
	x1, y1 := x0+g.tileset.TileWidth, y0+g.tileset.TileHeight

	return g.atlas.SubImage(image.Rect(x0, y0, x1, y1)).(*ebiten.Image)
}

// cursorPositionInTileSpace returns the coordinates of the tile the cursor is currently over
func (g *Game) cursorPositionInTileSpace() (int, int) {
	cx, cy := ebiten.CursorPosition()
	x := cx / (screenWidth / (g.myMap.MapWidth * g.tileset.TileWidth) * g.myMap.TileWidth)
	y := cy / (screenHeight / (g.myMap.MapHeight * g.tileset.TileHeight) * g.myMap.TileHeight)
	return x, y
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Game (Demo)")

	var atlas *ebiten.Image
	{
		imgFile, err := embeddedFS.Open("overworld.png")
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		img, _, err := image.Decode(imgFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		atlas = ebiten.NewImageFromImage(img)
	}

	// You can also read an image from your regular filesystem:
	// tiles, _, err := ebitenutil.NewImageFromFile("overworld.png")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(2)
	// }

	myMap, err := ebitmx.GetEbitenMapFromFS(embeddedFS, "map.tmx")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// You can also read a map from your regular filesystem:
	// myMap, err = ebitmx.GetEbitenMap("map.tmx")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(2)
	// }

	tileset, err := ebitmx.GetTilesetFromFS(embeddedFS, "tileset.tsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// You can also read a tileset from your regular filesystem:
	// tileset, err = ebitmx.GetEbitenTileset("map.tmx")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(2)
	// }

	game := &Game{
		myMap:   myMap,
		tileset: tileset,
		atlas:   atlas,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
