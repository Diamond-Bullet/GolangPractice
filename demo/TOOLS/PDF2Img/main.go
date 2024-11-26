package main

import (
	"GolangPractice/lib/logger"
	"fmt"
	"github.com/gen2brain/go-fitz"
	"github.com/gookit/color"
	"image/jpeg"
	"os"
)

const (
	Input        = "/tmp/input.pdf"
	OutputPrefix = "/tmp/output_"
)

func main() {
	// Open the PDF file
	doc, err := fitz.New(Input)
	if err != nil {
		logger.Error(err)
		return
	}
	defer doc.Close()

	// Iterate through the pages
	for i := 0; i < doc.NumPage(); i++ {
		img, localErr := doc.Image(i)
		if localErr != nil {
			logger.Error(localErr)
			return
		}

		// Create output file
		outFile, localErr := os.Create(OutputPrefix + fmt.Sprintf("%d.jpg", i+1))
		if localErr != nil {
			logger.Error(localErr)
			return
		}
		defer outFile.Close()

		// Encode as JPEG
		localErr = jpeg.Encode(outFile, img, &jpeg.Options{Quality: 90})
		if localErr != nil {
			logger.Error(localErr)
			return
		}

		color.Blueln("Page %d converted to output_%d.jpg\n", i+1, i+1)
	}
}
