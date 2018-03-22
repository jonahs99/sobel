package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"

	"github.com/jonahs99/sobel"
)

func main() {

	inPath := flag.String("in", "none", "path to input image")
	outPath := flag.String("out", "out.png", "path to output image")
	flag.Parse()

	// Reading the file as per stackoverflow.com/questions/8697095

	inFile, errInFile := os.Open(*inPath)
	outFile, errOutFile := os.Create(*outPath)

	if errInFile != nil {
		fmt.Printf("Error opening the input file '%s'.\n", *inPath)
		os.Exit(-1)
	}

	if errOutFile != nil {
		fmt.Printf("Error creating the output file '%s'.\n", *outPath)
		os.Exit(-1)
	}

	fmt.Println("Input file loaded.")

	src, _, errDecode := image.Decode(inFile)

	if errDecode != nil {
		fmt.Printf("Error decoding the input file.\n")
		os.Exit(-1)
	}

	fmt.Println("Input file decoded.")

	out := sobel.ApplySobel(src)

	fmt.Println("Sobel kernel applied.")

	defer outFile.Close()
	png.Encode(outFile, out)

	fmt.Printf("Output saved to '%s'.\n", *outPath)

}
