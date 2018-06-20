package main

import (
	"image"
	"log"
	"os"
	"fmt"
	"image/png"
	"flag"
	_ "image/jpeg"

	"github.com/urbanski/406hd/img"
	"github.com/disintegration/gift"
)

var (
	inputfile = flag.String("i","","input file")
	outputfile = flag.String("o","", "output file")
	filter = flag.String("f","", "filter name")
)

func IsValidFilter(filterName string) bool {
	switch filterName {
	case
		"twomedicine",
		"manyglacier",
		"stripheader":
			return true
	}
	return false
}

func main() {
	fmt.Println("406hd")

	flag.Parse()

	if *inputfile == "" || *outputfile == "" {
		fmt.Println("You must specify an input file (-i) and output file (-o)")
	}

	if !IsValidFilter(*filter) {
		fmt.Println(fmt.Sprintf("Invalid filter: '%s'", *filter))
		os.Exit(1)
	}

	if _, err := os.Stat(*inputfile); os.IsNotExist(err) {
		fmt.Println(fmt.Sprintf("Could not load '%s'", *inputfile))
		os.Exit(1)
	}

	src := loadImage(*inputfile)

	var g gift.Filter
	if *filter == "twomedicine" {
		g = img.TwoMedicine()
	} else if *filter == "manyglacier" {
		g = img.ManyGlacier()
	} else if *filter == "stripheader" {
		g = img.StripHeader()
	}


	dst := image.NewNRGBA(g.Bounds(src.Bounds()))


	g.Draw(dst, src, nil)
	saveImage(*outputfile, dst)
}


func loadImage(filename string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("os.Open failed: %v", err)
	}
	imgDecoded, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("image.Decode failed: %v", err)
	}
	return imgDecoded
}


func saveImage(filename string, img image.Image) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("os.Create failed: %v", err)
	}
	err = png.Encode(f, img)
	if err != nil {
		log.Fatalf("png.Encode failed: %v", err)
	}
}