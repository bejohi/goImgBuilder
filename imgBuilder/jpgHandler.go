package imgBuilder

import (
	"image"
	"image/jpeg"
	"os"
)

func loadJpg(filePath string) (image.Image, error){
	imgFile, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	return jpeg.Decode(imgFile)
}