package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"log"
	"os"

	"github.com/disintegration/imaging"
)

func main() {

	inputFile := flag.String("input", "", "Input image filename")
	outputFile := flag.String("output", "", "Output image filename")
	resizeWidth := flag.Int("resize_width", 0, "Width for resizing (optional)")
	resizeHeight := flag.Int("resize_height", 0, "Height for resizing (optional)")
	cropWidth := flag.Int("crop_width", 0, "Width for cropping (optional)")
	cropHeight := flag.Int("crop_height", 0, "Height for cropping (optional)")
	cropAnchor := flag.String("crop_anchor", "center", "Anchor point for cropping: top-left, top, top-right, left, center, right, bottom-left, bottom, or bottom-right (optional)")
	flag.Parse()

	if *inputFile == "" {
		fmt.Println("Input image filename is required")
		os.Exit(1)
	}

	src, err := imaging.Open(*inputFile)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	if *resizeWidth > 0 || *resizeHeight > 0 {
		src = imaging.Resize(src, *resizeWidth, *resizeHeight, imaging.Lanczos)
	}

	if *cropWidth > 0 && *cropHeight > 0 {
		anchor := imaging.Center
		switch *cropAnchor {
		case "top-left":
			anchor = imaging.TopLeft
		case "top":
			anchor = imaging.Top
		case "top-right":
			anchor = imaging.TopRight
		case "left":
			anchor = imaging.Left
		case "center":
			anchor = imaging.Center
		case "right":
			anchor = imaging.Right
		case "bottom-left":
			anchor = imaging.BottomLeft
		case "bottom":
			anchor = imaging.Bottom
		case "bottom-right":
			anchor = imaging.BottomRight
		}
		src = imaging.CropAnchor(src, *cropWidth, *cropHeight, anchor)
	}

	out, err := os.Create(*outputFile)
	if err != nil {
		log.Fatalf("failed to create output image: %v", err)
	}
	defer out.Close()

	err = jpeg.Encode(out, src, nil)
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

	fmt.Println("Image processing complete. Output saved to", *outputFile)
}
