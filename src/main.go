package main

import (
	"log"

	"github.com/iancullinane/stripper/src/config"
	"github.com/iancullinane/stripper/src/tile"
	"github.com/iancullinane/stripper/src/utils"
	"gopkg.in/gographics/imagick.v2/imagick"
)

const COLOR_COL = 4

func main() {

	config := config.New()

	log.Println("Start image processing")
	imagick.Initialize()

	// Orignical IM I used for first spriresheet
	// convert crom_thin.png +gravity -crop 32x32 +repage +adjoin output/%03d.png

	// Schedule cleanup
	defer imagick.Terminate()

	// Make a mw for the original
	mw := imagick.NewMagickWand()
	err := mw.ReadImage(config.GetInputFolder())
	if err != nil {
		panic(err)
	}

	// Create a second mw to write the images

	// Get the max number of tiles
	// TODO::validate tiles size as divisble by 16 or 32
	tilesW, tilesH := utils.GetTilesHandW(32, mw)
	log.Printf("Image is %d wide and %d tall", tilesW, tilesH)

	// Initialize a ten length slice of empty slices
	tiles := make([][]*tile.Tile, tilesH)
	for i := 0; i < tilesH; i++ {
		tiles[i] = make([]*tile.Tile, tilesW)
	}
	for h := 0; h < tilesH; h++ {
		for w := 0; w < tilesW; w++ {
			tiles[h][w] = tile.New(mw.Clone(), w, h, 32)
		}
	}

	uniqueSets := 0
	printers := make([]*imagick.MagickWand, 0)
	//
	var lastTile *tile.Tile
	for h := 0; h < tilesH; h++ {
		for w := 0; w < tilesW; w++ {
			if lastTile == nil && !tiles[h][w].HasOneColor() {
				printers = append(printers, imagick.NewMagickWand())

				uniqueSets++
			}

			printers[uniqueSets-1].AddImage(tiles[h][w].GetFinalImage())
			lastTile = tiles[h][w]
		}
	}

	// empty := false
	// TODO::each color for sprite 4 col
	// TODO::each class has single color tiles in between
	// log.Print(tiles)
	printers[0].WriteImages(config.GetOutputFolder(), true)
}
