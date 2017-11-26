package test

import (
	"testing"
	"github.com/bejohi/goImgBuilder/builder"
)

func TestImgBuilder_BuildWithoutFrom(t *testing.T){
	// Arrange & Act
	_, err := builder.ImgBuilder{}.Build()

	// Assert
	if err == nil{
		t.Error("An error was expected.")
	}
}
