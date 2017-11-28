package test

import (
	"testing"
	"github.com/bejohi/goImgBuilder/imgBuilder"
)

func TestImgBuilder_BuildWithoutFrom(t *testing.T){
	// Arrange & Act
	_, err := imgBuilder.ImgBuilder{}.Build()

	// Assert
	if err == nil{
		t.Error("An error was expected.")
	}
}


func TestImgBuilder_BuildWWithFromJpeg_WithWrongPath(t *testing.T){
	// Arrange & Act
	img, err := imgBuilder.ImgBuilder{}.FromJpeg("/wrong").Build()

	// Assert
	if err == nil || img != nil{
		t.Error("An error was expected.")
	}
}