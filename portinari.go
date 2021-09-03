// Copyright 2021 John Murowaniecki. All rights reserved.

// Package Portinari provide a tool to resize and crop given images.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"

	_ "image/jpeg"
	_ "image/png"

	"github.com/disintegration/imaging"
)

const (
	COMMON_APP_DESCRIPTION = "Portinari is a tool to resize and crop given images."
	COMMON_COPYRIGHT_NOTES = "Copyright 2021 John Murowaniecki. All rights reserved."
	MAJOR, MINOR, REVISION = 0, 0, 1
)

type Config struct {
	Source string
	Target string
	Size   string

	Width  int
	Height int

	Resize bool
	Crop   bool
}

var config Config

// Show detailed application usage info.
func Usage(COMMON_ERROR_MESSAGING string) {
	fmt.Printf(
		"%s\n%s\n\n%s\n\nExample:\n  %s --source my_original_image.jpeg --target my_new_image.png --size 200,400\n\nUsage:\n",
		COMMON_APP_DESCRIPTION,
		COMMON_COPYRIGHT_NOTES,
		COMMON_ERROR_MESSAGING,
		os.Args[0],
	)
	flag.PrintDefaults()

	os.Exit(1)
}

func init() {
	const (
		defaultImageSource = ""
		defaultImageTarget = "image.png"
		defaultResolutions = "150,200"
		___, ____, ______  = "", "", ""
	)
	flag.StringVar(&config.Source, "source", defaultImageSource, "Source image to be resized and croped.")
	flag.StringVar(&config.Target, "target", defaultImageTarget, "Target for processed image.")
	flag.StringVar(&config.Size, "size"+___, defaultResolutions, "W,H resolution.")
	flag.StringVar(&config.Size, "s"+______, defaultResolutions, "W,H resolution. (shorthand)")

	flag.IntVar(&config.Width, "width", 0, "Width.")
	flag.IntVar(&config.Width, "w"+___, 0, "Width. (shorthand)")

	flag.IntVar(&config.Height, "height", 0, "Height.")
	flag.IntVar(&config.Height, "h"+____, 0, "Height. (shorthand)")

	flag.Parse()
}

type Picture struct {
	Img    image.Image
	Src    image.Image
	Pos    image.Point
	Width  int
	Height int
}

func (i Picture) Open(file string) Picture {
	contents, err := os.Open(file)
	if nil != err {
		log.Fatal("Error0 ", err)
	}

	infos, _, err := image.DecodeConfig(contents)
	if nil != err {
		log.Fatal("Error1 ", err)
	}
	contents.Close()
	i.Height, i.Width = infos.Height, infos.Width

	i.Src, err = imaging.Open(file)
	if nil != err {
		log.Fatal("Error2 ", err)
	}
	i.Img = imaging.New(config.Width, config.Height, color.NRGBA{0, 0, 0, 0})
	i.Pos = image.Pt((config.Width/2)-(i.Width*(20000/i.Height)/100)/2, 0)
	return i
}

func (i Picture) Save() Picture {
	i.Src = imaging.CropAnchor(i.Src, i.Width, i.Height, imaging.Center)
	i.Src = imaging.Resize(i.Src, 0, config.Height, imaging.Lanczos)
	i.Img = imaging.Paste(i.Img, i.Src, i.Pos)
	fail := imaging.Save(i.Img, config.Target, imaging.PNGCompressionLevel(png.BestCompression))
	if fail != nil {
		log.Fatalf("failed to save image: %v\n\n%s / %s", fail, config.Source, config.Target)
	}
	return i
}

func asInt(numbers string) int {
	i, e := strconv.Atoi(numbers)
	if e != nil {
		return 0
	}
	return i
}

func Check(condition bool, setup *int, with string) {
	if condition {
		*setup = asInt(with)
	}
}

func main() {
	if config.Source == "" {
		Usage("Must provide an image source!")
	}
	resolution := strings.Split(config.Size, ",")
	width, height := resolution[0], resolution[1]

	Check(config.Width == 00, &config.Width, width)
	Check(config.Height == 0, &config.Height, height)

	Picture{}.
		Open(config.Source).
		Save()
}
