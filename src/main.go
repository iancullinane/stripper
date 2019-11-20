package main

import (
	"log"
	"os"

	"github.com/iancullinane/stripper/src/utils"
	"gopkg.in/gographics/imagick.v2/imagick"
)

const COLOR_COL = 4

func main() {

	input := "./test_file/gauis.png"
	output := "./output/%d.png"

	log.Println("Start image processing")
	imagick.Initialize()

	// Orignical IM I used for first spriresheet
	// convert crom_thin.png +gravity -crop 32x32 +repage +adjoin output/%03d.png

	// Schedule cleanup
	defer imagick.Terminate()

	if len(os.Args) < 1 {
		input = os.Args[1]
		output = os.Args[2]
	}

	// Make a mw for the original
	mw := imagick.NewMagickWand()
	err := mw.ReadImage(input)
	if err != nil {
		panic(err)
	}

	// Create a second mw to write the images
	mw_output := imagick.NewMagickWand()

	// Get the max number of tiles
	// TODO::validate tiles size as divisble by 16 or 32
	tilesW, tilesH := utils.GetTilesHandW(32, mw)
	log.Printf("Image is %d wide and %d tall", tilesW, tilesH)

	// TODO::each color for sprite 4 col
	// TODO::each class has single color tiles in between

	for h := 0; h < tilesH; h++ {
		for w := 0; w < COLOR_COL; w++ {
			// Cut out a tile
			temp := mw.Clone()
			temp.CropImage(32, 32, w*32, h*32)

			// A one color tile has no animation
			if utils.CheckIfOneColor(temp) {
				continue
			}

			mw_output.AddImage(temp)
		}
	}

	// Proces spritesheet
	// for h := 0; h < tilesH; h++ {
	// 	for w := 0; w < tilesW; w++ {

	// 		// Cut out a tile
	// 		temp := mw.Clone()
	// 		temp.CropImage(32, 32, w*32, h*32)

	// 		if utils.CheckIfOneColor(temp) {
	// 			continue
	// 		}

	// 		mw_output.AddImage(temp)
	// 	}

	// }
	mw_output.WriteImages(output, true)
}
