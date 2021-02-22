# Ebitmx

Ebitmx is a super simple parser to help [render](https://doc.mapeditor.org/en/latest/reference/tmx-map-format/) TMX maps when using [Ebiten](https://github.com/hajimehoshi/ebiten) for your games.

Right now is super limited to XML and CSV data structures in the TMX file, with support only for `orthogonal`, `right-down`, non `infinite` maps.

Please do **not** use this library in a production environment, this is just done as an example/helper to people who want to use TMX maps with Ebiten. And it's being used in the creation of a game. However, feel free to open a ticket to request a feature or send a PR.

This library parses a TMX file and returns a struct like the following, with the basic fields for you to render your map inside the `Draw()` function of the Ebiten main loop.

```go
type EbitenMap struct {
	TileWidth  int      // The width of the tile
	TileHeight int      // The height of the tile
	MapHeight  int      // The number of tiles in Height
	MapWidth   int      // The number of tiles in Width
	Layers     [][]int  // Layers
}
```

## Quick Start

```go
// TODO
```

## Roadmap
* Tilesets
* Orientation
* Renderorder
* Infinite maps

## License

MIT License