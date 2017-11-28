package imgBuilder

import (
	"image"
	"image/color"
	"errors"
)

type ImgBuilder struct {
	loadFromPath string
	loadFromJpgPath string
	loadFromPngPath string
	fromRectangle *image.Rectangle
	fromSizeWitdh int
	fromSizeHeight int
	withColor *color.Color
	saveAt string
	image image.Image
	imgLoadError error
}

func (builder ImgBuilder)FromFile(path string) ImgBuilder{
	builder.loadFromPath = path
	return builder
}

func (builder ImgBuilder)FromJpeg(path string) ImgBuilder{
	builder.loadFromPath = path
	return builder
}

func (builder ImgBuilder)FromRectangle(rectangle *image.Rectangle) ImgBuilder{
	builder.fromRectangle = rectangle
	return builder
}

func (builder ImgBuilder)FromSize(width int, height int) ImgBuilder{
	builder.fromSizeWitdh = width
	builder.fromSizeHeight = height
	return builder
}

func (builder ImgBuilder)AsColoredImage(color *color.Color) ImgBuilder{
	builder.withColor = color
	return builder
}

func (builder ImgBuilder)SaveAsFile(path string) ImgBuilder{
	builder.saveAt = path
	return builder
}

func (builder ImgBuilder)Build() (image.Image,error) {

	if !builder.checkFromOk(){
		return nil, errors.New("To build an image properly exactly one 'from' operation is required.")
	}




	return nil, errors.New("Build of new Image failed.")
}

