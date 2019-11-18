// Port of http://members.shaw.ca/el.supremo/MagickWand/resize.htm to Go
package main

import (
	"log"

	"gopkg.in/gographics/imagick.v2/imagick"
)

// "gopkg.in/gographics/imagick.v2/imagick"

func main() {

	input := "./test_file/gauis.png"
	// output := "./output/%d.png"

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

	mw := imagick.NewMagickWand()
	// err = mw.ReadImage(input)
	// if err != nil {
	// 	panic(err)
	// }
	// pw := imagick.NewPixelWand()
	// pw.SetColor("white")
	// mw.CropImage(32, 32, 0, 0)

	// Create a 100x100 image with a default of white
	mw.ReadImage(input)
	mw_output := imagick.NewMagickWand()

	// Get a new pixel iterator
	// iterator := mw.NewPixelRegionIterator(32, 32, 0, 0)

	for y := 0; y < int(mw.GetImageHeight()) && y < 10; y++ {
		for x := 0; x < int(mw.GetImageWidth()) && x < 10; x++ {
			log.Printf("%d: %d", x, y)
			temp := mw.Clone()
			temp.CropImage(32, 32, 0, 0)
			mw_output.AddImage(temp)
		}

	}
	mw_output.WriteImages("./output/%03d.png", true)

	// if err = mw.WriteImages(output, false); err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Wrote: %s\n", output)
}
