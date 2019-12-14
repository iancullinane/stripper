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
	// printers := imagick.NewMagickWand()

	// Get the max number of tiles
	// TODO::validate tiles size as divisble by 16 or 32
	tilesW, tilesH := utils.GetTilesHandW(32, mw)
	log.Printf("Image is %d wide and %d tall", tilesW, tilesH)

	// Initialize a ten length slice of empty slices
	tiles := make([][]*tile.Tile, tilesH)
	for i := 0; i < tilesH; i++ {
		tiles[i] = make([]*tile.Tile, tilesW)
	}
	// empty := false
	// TODO::each color for sprite 4 col
	// TODO::each class has single color tiles in between
	for h := 0; h < tilesH; h++ {
		for w := 0; w < COLOR_COL; w++ {

			tiles[h][w] = tile.New(mw.Clone(), w, h, 32)
			// if temp.HasOneColor() {
			// 	continue
			// }

			// printers.AddImage(temp.GetFinalImage())
		}
	}
	log.Printf("%#v", tiles)
	// printers.WriteImages(config.GetOutputFolder(), true)
}
