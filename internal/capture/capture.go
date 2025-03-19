package capture

import (
	"fmt"
	"image"
	"log"

	"github.com/go-vgo/robotgo"
	"github.com/kbinani/screenshot"
)

// FullScreenUsingScreenshot captures the entire primary screen
// Returns an image.Image or an error.
func FullScreenUsingScreenshot() (image.Image, error) {
	// Get the number of active displays.
	displayCount := screenshot.NumActiveDisplays()
	if displayCount < 1 {
		return nil, fmt.Errorf("no active display found")
	}

	// We take the bounds of the primary display (index 0).
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return nil, fmt.Errorf("failed to capture screen: %w", err)
	}

	return img, nil
}

// FullScreenUsingRobotGo captures the entire screen using robotgo.
// Returns an image.Image or an error.
func FullScreenUsingRobotGo() (image.Image, error) {
	// robotgo captures the main screen by default, returns a bitmap in memory.
	bitMap := robotgo.CaptureScreen()
	if bitMap == nil {
		return nil, fmt.Errorf("failed to capture screen with robotgo")
	}
	defer robotgo.FreeBitmap(bitMap)

	// Convert robotgo C-struct to image.Image
	img := robotgo.ToImage(bitMap)
	if img == nil {
		return nil, fmt.Errorf("failed to convert robotgo bitmap to image")
	}

	return img, nil
}

// debugging function.
func DebugCapture() {
	img, err := FullScreenUsingScreenshot()
	if err != nil {
		log.Printf("Capture error (screenshot): %v\n", err)
		return
	}
	log.Println("Captured screen using kbinani/screenshot:", img.Bounds())

	img2, err := FullScreenUsingRobotGo()
	if err != nil {
		log.Printf("Capture error (robotgo): %v\n", err)
		return
	}
	log.Println("Captured screen using robotgo:", img2.Bounds())
}
