[![Go Report Card](https://goreportcard.com/badge/github.com/zeozeozeo/imagesim)](https://goreportcard.com/report/github.com/zeozeozeo/imagesim)
![GitHub](https://img.shields.io/github/license/zeozeozeo/imagesim)
# image SIMililarity

This is a fast and simple algorithm for comparing the similarity of 2 images. Pure Go and zero dependencies.

# [Example](https://github.com/zeozeozeo/imagesim/blob/main/internal/example/compare.go)

## The hashing algorithm

1. Resize the images to be 8 by 8 pixels. This is done by a very simple resizing algorithm.
2. Make the images grayscale.
3. Get the bitmask image value threshold. This is done by getting the average gray value in the image.
4. Calculate the bitmask of the image, skip all values that are smaller than the threshold.

## The hash comparing algorithm

1. Calculate the [Hamming distance](https://en.wikipedia.org/wiki/Hamming_distance) between two hashes.
2. Divide it by the amount of bits in the hash (this library uses uint64 for hashes, so 64).