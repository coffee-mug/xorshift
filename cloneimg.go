package main

import (
	"flag"
	"github.com/coffee-mug/xorshift/xorshift32"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {

	// Get in et output file names from cmdline
	in := flag.String("in", "in.png", "Source png path to be xored. Defaults to in.png")
	out := flag.String("out", "out.png", "Xored png path/name. Defaults to out.png")
	key := flag.Int("key", 100, "Uint32 key used to seed the PRNG")

	flag.Parse()

	inFile, err := os.Open(*in)
	defer inFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	inImage, err := png.Decode(inFile)
	if err != nil {
		log.Fatal(err)
	}

	outImage := image.NewNRGBA(inImage.Bounds())
	if err != nil {
		log.Fatal(err)
	}

	// init random source
	seed := uint32(*key)
	randomizer := xorshift32.Xorshift32(&seed)

	for y := inImage.Bounds().Min.Y; y < inImage.Bounds().Max.Y; y++ {
		for x := inImage.Bounds().Min.X; x < inImage.Bounds().Max.X; x++ {
			r, g, b, a := inImage.At(x, y).RGBA()
			outImage.Set(x, y, color.NRGBA{
				R: uint8(r ^ randomizer()),
				G: uint8(g ^ randomizer()),
				B: uint8(b ^ randomizer()),
				A: uint8(a),
			})
		}
	}

	f, err := os.Create(*out)
	if err != nil {
		log.Fatal()
	}
	defer f.Close()

	png.Encode(f, outImage)
}
