package xorshift32

import (
  "fmt"
  "image"
  "image/png"
  "os"
)

// Xorshift32 returns a 32bits random number
// State must be initialized ot non-zero
// the wikipedia C version uses a static variable
// we use a closure to mimic this.
// Marsaglia recommended numbers for uint32: 13, 17, 5
func Xorshift32(state *uint32) func() uint32 {
  var x uint32
  return func() uint32 {
    x = *state
    x = x ^ (x << 13)
    x = x ^ (x >> 17)
    x = x ^ (x << 5)
    *state = x
    return x
  }
}

// good numbers for uint16: 11, 8, 5 
func xorshift16(state *uint16) func() uint16 {
  var x uint16
  return func() uint16 {
    x = *state
    x = x ^ (x << 5)
    x = x ^ (x >> 3)
    x = x ^ (x << 13)
    *state = x
    return x - 1
  }
}

func xorshift128plus(fst, snd uint64) uint64 {
  var x, y uint64
  s := []uint64{fst,snd}
  x = s[0]
  y = s[1]

  s[0] = y

  x = x ^ (x << 23) // a
  s[1] = x ^ y ^ (x >> 17) ^ (y >> 26) // b, c
  return s[1] + y
}

func createImage() {
  myImage := image.NewNRGBA(image.Rect(0, 0, 300, 600))

  // Feed the pixels with random values
  // One pixel takes up four bytes/uint8. One for each RGBA

  // PRNG function
  st := uint32(10)
  xor := Xorshift32(&st)

  for i := 0; i < 180000; i++ {
    myImage.Pix[i] = uint8(xor())
  }

  outputFile, err := os.Create("test.png")
  defer outputFile.Close()

  if err != nil {
    fmt.Println("Something bad happened mate", err)
  }


  png.Encode(outputFile, myImage)
}

func main() {
  createImage()
}
