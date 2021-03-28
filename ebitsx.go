package ebitmx

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"math"
	"os"
)

// EbitenTileset is a friendly representation of a TSX Tileset
type EbitenTileset struct {
	TileWidth     int
	TileHeight    int
	TilesetWidth  int
	TilesetHeight int
	Tiles         []Tile
}

// GetEbitenTileset returns a simplified TSX Tileset, based on a file on disk
func GetEbitenTileset(path string) (*EbitenTileset, error) {
	return GetTilesetFromFS(os.DirFS("."), path)
}

// GetTilesetFromFS allows you to pass in the file system used to find the desired file
// This is useful for Go's v1.16 embed package which makes it simple to embed assets into
// your binary and accessible via the embed.FS which is compatible with the fs.FS interface
func GetTilesetFromFS(fileSystem fs.FS, path string) (*EbitenTileset, error) {
	tsxFile, err := fileSystem.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening TSX file %s: %v", path, err)
	}
	defer tsxFile.Close()

	bytes, err := ioutil.ReadAll(tsxFile)
	if err != nil {
		return nil, fmt.Errorf("error reading TSX file %s: %v", path, err)
	}

	tsxTileset, err := ParseTSX(bytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing TSX file %s: %v", path, err)
	}

	return transformMapToEbitenTileset(tsxTileset)
}

func transformMapToEbitenTileset(tsx *Tileset) (*EbitenTileset, error) {
	ebitenMap := &EbitenTileset{
		TileWidth:     tsx.TileWidth,
		TileHeight:    tsx.TileHeight,
		TilesetWidth:  tsx.Columns,
		TilesetHeight: int(math.Ceil(float64(tsx.TileCount) / float64(tsx.Columns))),
		Tiles:         tsx.Tiles,
	}

	return ebitenMap, nil
}
