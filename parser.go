package ebitmx

import (
	"encoding/xml"
	"fmt"
)

// ParseTMX parses a TMX file and returns a Map
// For now, we only allow XML format
func ParseTMX(bytes []byte) (*Map, error) {
	tmxMap := &Map{}
	err := xml.Unmarshal(bytes, tmxMap)
	if err != nil {
		return nil, fmt.Errorf("only <xml> format is allowed: %v", err)
	}

	err = checkLimitations(tmxMap)
	if err != nil {
		return nil, fmt.Errorf("unsupported: %v", err)
	}
	return tmxMap, nil
}

// TODO address limitations
func checkLimitations(tmxMap *Map) error {
	if tmxMap.Orientation != "orthogonal" {
		return fmt.Errorf("orientation must be orthogonal")
	}

	if tmxMap.RenderOrder != "right-down" {
		return fmt.Errorf("renderorder must be right-down")
	}

	if tmxMap.Infinite {
		return fmt.Errorf("infinite is not supported")
	}

	return nil
}
