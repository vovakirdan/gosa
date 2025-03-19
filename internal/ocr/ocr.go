package ocr

import (
	"bytes"
	"fmt"
	"image"
	"image/png"

	"github.com/otiai10/gosseract"
)

// ScanResult contains recognized text and bounding box coordinates.
type ScanResult struct {
	Text   string
	Bounds image.Rectangle
}

// RecognizeTextWithCoordinates performs OCR on a grayscale image
// and returns slices of text fragments with bounding boxes.
func RecognizeTextWithCoordinates(img image.Image) ([]ScanResult, error) {
	// 1) Преобразовать в градации серого
	grayImg := GrayscaleImage(img)

	// 2) Кодируем изображение в PNG (in-memory)
	var buf bytes.Buffer
	if err := png.Encode(&buf, grayImg); err != nil {
		return nil, fmt.Errorf("failed to encode grayscale image: %w", err)
	}

	// 3) Инициализируем gosseract
	client := gosseract.NewClient()
	defer client.Close()

	// Указываем язык (по желанию)
	err := client.SetLanguage("eng")
	if err != nil {
		return nil, fmt.Errorf("failed to set language: %w", err)
	}

	// Передаём байты в tesseract
	if err := client.SetImageFromBytes(buf.Bytes()); err != nil {
		return nil, fmt.Errorf("failed to set image bytes: %w", err)
	}

	// 4) Получаем список bounding boxes
	boxes, err := client.GetBoundingBoxes(gosseract.RIL_WORD)
	if err != nil {
		return nil, fmt.Errorf("failed to get bounding boxes: %w", err)
	}

	// Формируем результат
	var results []ScanResult
	for _, box := range boxes {
		// box.Word содержит распознанное слово
		// box.Box содержит {X, Y, W, H}
		// Преобразуем в image.Rectangle
		rect := image.Rect(box.Box.Min.X, box.Box.Min.Y, box.Box.Min.X+box.Box.Size().X, box.Box.Min.Y+box.Box.Size().Y)

		results = append(results, ScanResult{
			Text:   box.Word,
			Bounds: rect,
		})
	}

	return results, nil
}
