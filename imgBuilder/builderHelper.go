package imgBuilder

func (builder ImgBuilder)checkFromOk() bool {
	fromCounter := 0

	if builder.fromSizeHeight != 0{
		fromCounter++
	}

	if builder.fromRectangle != nil {
		fromCounter++
	}

	if builder.loadFromPath != ""{
		fromCounter++
	}

	if builder.loadFromJpgPath != ""{
		fromCounter++
	}

	if builder.loadFromPngPath != ""{
		fromCounter++
	}

	if fromCounter != 1{
		return false
	}

	return true
}

func (builder ImgBuilder)handleFromOperation(){
	if builder.loadFromJpgPath != ""{
		builder.image, builder.imgLoadError = loadJpg(builder.loadFromJpgPath)
	}
}