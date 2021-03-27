package ebitmx

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

// GetEbitenMap returns a map that Ebiten can understand
// based on a TMX file. Note that some data might be lost, as Ebiten
// does not require too much information to render a map
// func GetEbitenMap(path string) (*EbitenMap, error) {
// 	return GetEbitenMapFromFS(os.DirFS("."), path)
// }

// func GetTilesetFromFS(fileSystem fs.FS, path string) (*EbitenMap, error) {
func GetTilesetFromFS(fileSystem fs.FS, path string) (*Tileset, error) {
	tsxFile, err := fileSystem.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Error opening TSX file %s: %v", path, err)
	}

	defer tsxFile.Close()

	bytes, err := ioutil.ReadAll(tsxFile)
	if err != nil {
		return nil, fmt.Errorf("Error reading TSX file %s: %v", path, err)
	}

	tsxTileset, err := ParseTSX(bytes)
	if err != nil {
		return nil, fmt.Errorf("Error parsing TSX file %s: %v", path, err)
	}

	return tsxTileset, nil

	// return transformMapToEbitenMap(tmxMap)
}

// func transformMapToEbitenMap(tmx *Map) (*EbitenMap, error) {
// 	ebitenMap := &EbitenMap{
// 		TileWidth:  tmx.TilHeight,
// 		TileHeight: tmx.TileWidth,
// 		MapHeight:  tmx.Height,
// 		MapWidth:   tmx.Width,
// 	}

// 	var ebitenLayers [][]int
// 	for _, layer := range tmx.Layers {
// 		var innerLayer []int
// 		if layer.Data.Encoding == "csv" {

// 			for _, s := range strings.Split(string(layer.Data.Raw), ",") {
// 				s = strings.TrimSpace(s)
// 				coord, err := strconv.Atoi(s)

// 				if err != nil {
// 					return nil, fmt.Errorf("Error parsing layer [%s] data, %v is not a number", layer.Name, s)
// 				}
// 				innerLayer = append(innerLayer, coord)
// 			}

// 		}
// 		ebitenLayers = append(ebitenLayers, innerLayer)
// 	}

// 	ebitenMap.Layers = ebitenLayers
// 	return ebitenMap, nil
// }
