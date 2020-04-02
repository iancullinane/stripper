package main

import (
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/iancullinane/stripper/src/config"
	"github.com/iancullinane/stripper/src/tile"
	"github.com/iancullinane/stripper/src/utils"
	"gopkg.in/gographics/imagick.v2/imagick"
)

const COLOR_COL = 4

func main() {

	config := config.New()

	log.Println("Start image processing")
	od := config.GetOutputFolder()

	err := ClearDir(od)
	if err != nil {
		panic(err)
	}

	imagick.Initialize()

	// Orignical IM I used for first spriresheet
	// convert crom_thin.png +gravity -crop 32x32 +repage +adjoin output/%03d.png

	// Schedule cleanup
	defer imagick.Terminate()

	// Make a mw for the original
	og := imagick.NewMagickWand()
	err = og.ReadImage(config.GetInputFolder())
	if err != nil {
		panic(err)
	}

	// Create a second mw to write the images

	// Get the max number of tiles
	// TODO::validate tiles size as divisble by 16 or 32
	tilesW, tilesH := utils.GetTilesHandW(32, og)
	log.Printf("Image is %d wide and %d tall", tilesW, tilesH)

	// Initialize a ten length slice of empty slices
	tiles := make([][]*tile.Tile, tilesH)
	for i := 0; i < tilesH; i++ {
		tiles[i] = make([]*tile.Tile, tilesW)
	}
	for h := 0; h < tilesH; h++ {
		for w := 0; w < tilesW; w++ {
			tiles[h][w] = tile.New(og.Clone(), w, h, 32)
		}
	}

	uniqueSets := 0
	printers := make([]*imagick.MagickWand, 0)

	// var lastTile *tile.Tile
	for h := 0; h < tilesH; h++ {
		for w := 0; w < tilesW; w++ {

			if tiles[h][w].HasOneColor() {
				printers = append(printers, imagick.NewMagickWand())
				uniqueSets++
				continue
			}

			printers = append(printers, imagick.NewMagickWand())
			if uniqueSets == 0 {
				printers[uniqueSets].AddImage(tiles[h][w].GetFinalImage())
			} else {
				printers[uniqueSets-1].AddImage(tiles[h][w].GetFinalImage())
			}
			// lastTile = tiles[h][w]
		}
	}

	log.Println(config.GetOutputFolder())
	log.Println(len(printers))
	// empty := false
	// TODO::each color for sprite 4 col
	// TODO::each class has single color tiles in between
	// log.Print(tiles)
	for _, printer := range printers {
		printer.WriteImages(config.GetOutputFolder(), true)
	}

}

func ClearDir(dir string) error {

	cleanPath := path.Dir(dir)
	log.Printf("Remove %s", cleanPath)
	d, err := os.Open(cleanPath)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		log.Printf("rm %s/%s", cleanPath, name)
		err = os.RemoveAll(filepath.Join(cleanPath, name))
		if err != nil {
			return err
		}
	}
	return nil
}
