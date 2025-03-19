package main

import (
	"fmt"
	"log"

	"github.com/vovakirdan/gosa/internal/capture"
	"github.com/vovakirdan/gosa/internal/ocr"
)

func main() {
	// 1) Захват экрана (к примеру, kbinani/screenshot)
	img, err := capture.FullScreenUsingScreenshot()
	if err != nil {
		log.Fatalf("Error capturing screen: %v", err)
	}

	// 2) Распознаём текст и координаты
	results, err := ocr.RecognizeTextWithCoordinates(img)
	if err != nil {
		log.Fatalf("Error recognizing text: %v", err)
	}

	// 3) Выводим результаты в консоль
	for i, r := range results {
		fmt.Printf("Fragment #%d: \"%s\" at %v\n", i+1, r.Text, r.Bounds)
	}
}
