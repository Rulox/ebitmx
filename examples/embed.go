package main

import "embed"

// embeddedFS  holds our game assets
//go:embed map.tmx overworld.png
var embeddedFS embed.FS
