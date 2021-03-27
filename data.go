package ebitmx

import "encoding/xml"

// Tileset represents a set of tiles in a TMX, or a TSX file
type Tileset struct {
	XMLName    xml.Name `xml:"tileset"`
	Version    string   `xml:"version,attr"`
	Name       string   `xml:"name,attr"`
	TileWidth  int      `xml:"tilewidth,attr"`
	TileHeight int      `xml:"tileheight,attr"`
	TileCount  int      `xml:"tilecount,attr"`
	Columns    int      `xml:"columns,attr"`
	Image      Image    `xml:"image"`
}

type Image struct {
	XMLName xml.Name `xml:"image"`
	Format  string   `xml:"format"`
	Source  string   `xml:"source"`
	Width   int      `xml:"width,attr"`
	Height  int      `xml:"height,attr"`
}

// Map is the representation of a map in a TMX file
type Map struct {
	XMLName      xml.Name `xml:"map"`
	Version      string   `xml:"version,attr"`
	TiledVersion string   `xml:"tiledversion,attr"`
	Orientation  string   `xml:"orientation,attr"`
	RenderOrder  string   `xml:"renderorder,attr"`
	Width        int      `xml:"width,attr"`
	Height       int      `xml:"height,attr"`
	TileWidth    int      `xml:"tilewidth,attr"`
	TilHeight    int      `xml:"tileheight,attr"`
	Infinite     bool     `xml:"infinite,attr"`
	// TODO nextlayerid and nextobjectid ?

	//Tileset []Tileset `xml:"tileset"`
	Layers []Layer `xml:"layer"`
}

// Layer represents a layer in the TMX map file
type Layer struct {
	XMLName xml.Name `xml:"layer"`
	ID      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
	Data    Data     `xml:"data"`
	Width   int      `xml:"width,attr"`
	Height  int      `xml:"height,attr"`
}

// Data represents the data inside a Layer
type Data struct {
	XMLName  xml.Name `xml:"data"`
	Encoding string   `xml:"encoding,attr"`
	Raw      []byte   `xml:",innerxml"`
}

// EbitenMap is the transformed representation of a TMX map in the simplest
// way possible for Ebiten to understand and render
type EbitenMap struct {
	TileWidth  int
	TileHeight int
	MapHeight  int
	MapWidth   int
	Layers     [][]int
}
