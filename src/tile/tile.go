package tile

import (
	"log"

	"gopkg.in/gographics/imagick.v2/imagick"
)

type Tile struct {
	x   int
	y   int
	img *imagick.MagickWand
}

func New(img *imagick.MagickWand, x, y int, size uint) *Tile {

	err := img.CropImage(size, size, x*int(size), y*int(size))
	if err != nil {
		log.Printf("Tile creation failed at [%d][%d]: %s", x, y, err)
	}

	return &Tile{
		x:   x,
		y:   y,
		img: img,
	}
}

func (t Tile) HasOneColor() bool {

	uc := t.img.GetImageColors()
	if uc == 1 {
		return true
	}
	return false
}

func (t Tile) GetFinalImage() *imagick.MagickWand {
	return t.img
}

// func NewProcessor() *tile {
// 	return &tile{
// 		prev *imagick.MagickWand
// 		curr *imagick.MagickWand
// 		next *imagick.MagickWand
// 	}
// }
