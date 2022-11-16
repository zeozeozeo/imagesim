package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image/jpeg"
	"time"

	"github.com/zeozeozeo/imagesim"
)

//go:embed parrot1.jpg
var parrot1Data []byte // by https://unsplash.com/@kriztheman

//go:embed parrot2.jpg
var parrot2Data []byte // by https://unsplash.com/@ideasyormanch

func main() {
	// decode the images
	fmt.Println("decoding images...")
	start := time.Now()

	parrot1, err := jpeg.Decode(bytes.NewBuffer(parrot1Data))
	if err != nil {
		panic(err)
	}
	parrot2, err := jpeg.Decode(bytes.NewBuffer(parrot2Data))
	if err != nil {
		panic(err)
	}
	fmt.Printf("decoded in %s\n\n", time.Since(start))
	start = time.Now()

	// now, calculate the image hashes and compare them

	// store this hash for later use if you want to compare multiple images
	parrot2Hash := imagesim.Hash(parrot2)

	// compare the image and the hash of the other image
	// you can use CompareHashes for comparing 2 hashes, or CompareImages
	// for comparing two images (will generate 2 hashes each time).
	diff := imagesim.Compare(parrot1, parrot2Hash)

	fmt.Printf("calculated 2 hashes and their difference in %s\n\n", time.Since(start))
	fmt.Printf("difference: %f\n", diff) // "difference: 0.531250"
}
