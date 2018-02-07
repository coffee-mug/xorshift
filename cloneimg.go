package main

import (
  "image"
  "image/png"
  "image/color"
  "os"
  "log"
  "github.com/xoroshiro128plus/xorshift32"
)

func main() {
  inFile, err := os.Open("out.png")
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
  seed := uint32(32)
  randomizer := xorshift32.Xorshift32(&seed)

  for y:= inImage.Bounds().Min.Y; y < inImage.Bounds().Max.Y; y++ {
    for x := inImage.Bounds().Min.X; x < inImage.Bounds().Max.X; x++ {
      r, g, b, a := inImage.At(x,y).RGBA()
      outImage.Set(x, y, color.NRGBA{
        R: uint8(r ^ randomizer()),
        G: uint8(g ^ randomizer()),
        B: uint8(b ^ randomizer()),
        A: uint8(a),
      })
    }
  }

  f, err := os.Create("in2.png")
  if err != nil {
    log.Fatal()
  }
  defer f.Close()

  png.Encode(f, outImage)
}
