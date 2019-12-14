package utils

import (
	"gopkg.in/gographics/imagick.v2/imagick"
)

func GetTilesHandW(tileSize int, img *imagick.MagickWand) (w, h int) {
	tilesW := int(img.GetImageWidth()) / tileSize
	tilesH := int(img.GetImageHeight()) / tileSize

	return tilesW, tilesH
}

func CropTile(img imagick.MagickWand, xPos, yPos int) *imagick.MagickWand {
	// Cut out a tile
	temp := img.Clone()
	temp.CropImage(32, 32, xPos*32, yPos*32)

	return temp
}

func CheckIfOneColor(img *imagick.MagickWand) bool {
	uc := img.GetImageColors()
	if uc == 1 {
		return true
	}
	return false
}
