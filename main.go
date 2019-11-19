// Port of http://members.shaw.ca/el.supremo/MagickWand/resize.htm to Go
package main

import (
	"log"

	"gopkg.in/gographics/imagick.v2/imagick"
)

func main() {

	input := "./test_file/gauis.png"
	output := "./output/%d.png"

	log.Println("Start image processing")
	imagick.Initialize()

	// convert crom_thin.png +gravity -crop 32x32 +repage +adjoin output/%03d.png

	// Schedule cleanup
	defer imagick.Terminate()
	// var err error

	// if len(os.Args) < 1 {
	// 	input = os.Args[1]
	// 	output = os.Args[2]
	// }

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
	tilesW := int(mw.GetImageWidth()) / 32
	tilesH := int(mw.GetImageHeight()) / 32
	log.Printf("Image is %d wide and %d tall", tilesW, tilesH)

	// Proces spritesheet
	for h := 0; h < tilesH; h++ {
		for w := 0; w < tilesW; w++ {

			// Cut out a tile
			temp := mw.Clone()
			temp.CropImage(32, 32, w*32, h*32)

			uc := temp.GetImageColors()
			log.Printf("%d", uc)

			if uc == 1 {
				log.Printf("Frame with only one color at %d: %d", h, w)
				continue
			}

			mw_output.AddImage(temp)
		}

	}
	mw_output.WriteImages(output, true)
}
