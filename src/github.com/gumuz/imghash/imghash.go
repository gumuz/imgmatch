package imghash

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func MeanHash(path string) (uint64, error) {
	var (
		img         image.Image
		total, grey uint32
		pixels      []uint32
		result      uint64
	)

	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	switch {
	case strings.Contains(path, ".png"):
		img, err = png.Decode(file)
	case strings.Contains(path, ".jpg"):
		img, err = jpeg.Decode(file)
	}

	if err != nil {
		return 0, err
	}

	// first pass, averaging colors from resized src image
	wRatio := img.Bounds().Size().X / 8
	hRatio := img.Bounds().Size().Y / 8
	for i := 0; i < 64; i++ {
		c := img.At((i%8)*wRatio, (i/8)*hRatio)
		r, g, b, a := c.RGBA()
		r, g, b = r/256, g/256, b/256
		if a == 0 {
			r, g, b = 255, 255, 255
		}
		grey = (r + g + b) / 3
		total += grey
		pixels = append(pixels, grey)
	}

	mean := total / 64
	//create hash
	for idx, c := range pixels {
		if uint32(c) > mean {
			result = result | 1<<uint(idx)
		}
	}

	return result, nil
}
