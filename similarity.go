package imagesim

import (
	"image"
)

// Hash returns the comparable hash of the image.
func Hash(img image.Image) uint64 {
	// TODO: this always downscales the image to 8x8 pixels, add a way to specify that
	grayImg := desaturateAndResizeImage(img, 8, 8)
	return calcImageBitmask(grayImg, getImageAverage(grayImg))
}

// Compare compares an image with an image hash.
func Compare(img image.Image, hash uint64) float64 {
	return CompareHashes(Hash(img), hash)
}

// CompareHashes compares two hashes and returns their similarity (lower = more similar).
func CompareHashes(hash1, hash2 uint64) float64 {
	distance := hammingDistance(hash1, hash2)
	return float64(distance) / 64
}

// CompareImages compares two images and returns the similarity (lower = more similar)
// between them. You shouldn't use this because it will calculate the hash for both images
// at the same time, instead you should only calculate the hash once by calling Hash(img),
// and then passing the image you want to compare with the hash of the other image to Compare.
func CompareImages(img1, img2 image.Image) float64 {
	return CompareHashes(Hash(img1), Hash(img2))
}

// desaturateAndResizeImage makes the image grayscale and resizes it to the specified
// width and height.
func desaturateAndResizeImage(img image.Image, width, height int) *image.Gray {
	// scaling factor
	scaleX := float64(img.Bounds().Dx()) / float64(width)
	scaleY := float64(img.Bounds().Dy()) / float64(height)
	grayImg := image.NewGray(image.Rect(0, 0, width, height))

	// naive image scaling algorithm
	for nx := 0; nx < width; nx++ {
		for ny := 0; ny < height; ny++ {
			sx := int(float64(nx) * scaleX)
			sy := int(float64(ny) * scaleY)
			grayImg.Set(nx, ny, img.At(sx, sy))
		}
	}
	return grayImg
}

// getImageAverage returns the average pixel value in a grayscale image.
func getImageAverage(img *image.Gray) int {
	total := 0
	for _, val := range img.Pix {
		total += int(val)
	}
	return total / len(img.Pix)
}

// calcImageBitmask generates a comparable image bitmask.
func calcImageBitmask(img *image.Gray, threshold int) uint64 {
	result := uint64(0)
	for idx, val := range img.Pix {
		if int(val) >= threshold {
			result |= 1 << idx
		}
	}
	return result
}

// hammingDistance calculates the Hamming distance between two series of bits.
func hammingDistance(x, y uint64) (dist uint64) {
	x ^= y
	for x > 0 {
		dist += 1
		x &= x - 1
	}
	return
}
