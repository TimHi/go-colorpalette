package src

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
)

func GetColors(url string) ([]string, error) {
	// Just a simple GET request to the image URL
	// We get back a *Response, and an error
	res, err := http.Get(url)

	if err != nil {
		log.Fatalf("http.Get -> %v", err)
	}

	// We read all the bytes of the image
	// Types: data []byte
	data, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("ioutil.ReadAll -> %v", err)
	}

	// You have to manually close the body, check docs
	// This is required if you want to use things like
	// Keep-Alive and other HTTP sorcery.
	res.Body.Close()

	// Decode the image based on its format (JPEG or PNG)
	img, format, err := image.Decode(bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return []string{}, err
	}
	colors := []color.Color{}
	fmt.Println("Image format:", format)
	fmt.Println(img.At(1, 1))
	return []string{}, nil
}
