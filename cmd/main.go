package main

import (
	"fmt"
	"log"
	"os"
	"image/png"

	"github.com/vovakirdan/gosa/internal/capture"
	"github.com/vovakirdan/gosa/internal/ocr"
)

func main() {
	// 1) Захватываем все экраны
	images, err := capture.CaptureAllDisplaysUsingScreenshot()
	if err != nil {
		log.Fatalf("Error capturing screens: %v", err)
	}

	// Перебираем каждый дисплей
	for displayIndex, img := range images {
		// Сохраняем для отладки
        outFile, err := os.Create(fmt.Sprintf("screen_display_%d.png", displayIndex))
        if err != nil {
            log.Printf("Error creating file for display %d: %v", displayIndex, err)
            continue
        }
        if err := png.Encode(outFile, img); err != nil {
            log.Printf("Error encoding PNG for display %d: %v", displayIndex, err)
        }
        outFile.Close()
		// 2) Распознаём текст и координаты (только английский язык)
		results, err := ocr.RecognizeTextWithCoordinates(img)
		if err != nil {
			log.Printf("Error recognizing text on display %d: %v", displayIndex, err)
			continue
		}

		// 3) Выводим результаты в консоль
		fmt.Printf("=== Display %d: found %d text fragments ===\n", displayIndex, len(results))
		for i, r := range results {
			fmt.Printf("  Fragment #%d: \"%s\" at %v\n", i+1, r.Text, r.Bounds)
		}
	}
}
