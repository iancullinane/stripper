package utils

import (
	"gopkg.in/gographics/imagick.v2/imagick"
)

func GetTilesHandW(tileSize int, img *imagick.MagickWand) (w, h int) {
	tilesW := int(img.GetImageWidth()) / tileSize
	tilesH := int(img.GetImageHeight()) / tileSize

	return tilesW, tilesH
}

func CheckIfOneColor(img *imagick.MagickWand) bool {
	uc := img.GetImageColors()
	if uc == 1 {
		return true
	}
	return false
}
