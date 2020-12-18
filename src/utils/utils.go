package utils

import (
	"gopkg.in/gographics/imagick.v2/imagick"
)

// GetTilesHandW will return the number of tiles in the provided spritesheet
func GetTilesHandW(tileSize int, img *imagick.MagickWand) (w, h int) {
	tilesW := int(img.GetImageWidth()) / tileSize
	tilesH := int(img.GetImageHeight()) / tileSize

	return tilesW, tilesH
}

// CropTile cuts the tile out of the image
func CropTile(img imagick.MagickWand, xPos, yPos int) *imagick.MagickWand {
	// Cut out a tile
	temp := img.Clone()
	temp.CropImage(32, 32, xPos*32, yPos*32)

	return temp
}

// CheckIfOneColor returns true if the tile has only one color
func CheckIfOneColor(img *imagick.MagickWand) bool {
	uc := img.GetImageColors()
	if uc == 1 {
		return true
	}
	return false
}
