package ocr

import (
	"image"
	"image/draw"
)

// GrayscaleImage converts the given image to grayscale.
// This helps improve OCR accuracy in many cases.
func GrayscaleImage(src image.Image) *image.Gray {
	bounds := src.Bounds()
	grayImg := image.NewGray(bounds)

	// Draw src onto the grayImg
	draw.Draw(grayImg, bounds, src, bounds.Min, draw.Src)

	return grayImg
}
