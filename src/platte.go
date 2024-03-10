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
	"sort"
)

type colorCount struct {
	color color.Color
	count int
}

func GetColors(url string) ([]color.Color, error) {
	res, err := http.Get(url)

	if err != nil {
		log.Printf("http.Get -> %v \n", err)
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		log.Println(err)
	}
	res.Body.Close()

	img, _, err := image.Decode(bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return []color.Color{}, err
	}

	colorCounts := make(map[color.Color]int)
	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			colorCounts[img.At(x, y)]++
		}
	}
	mostProminentColors := getMostProminentColors(colorCounts)

	rgbColors := []color.Color{}
	for _, color := range mostProminentColors {
		rgbColors = append(rgbColors, color.color)
	}
	return rgbColors, nil
}

func getMostProminentColors(colorCounts map[color.Color]int) []colorCount {
	var colorCountSlice []colorCount
	for c, count := range colorCounts {
		colorCountSlice = append(colorCountSlice, struct {
			color color.Color
			count int
		}{color: c, count: count})
	}

	sort.Slice(colorCountSlice, func(i, j int) bool {
		return colorCountSlice[i].count > colorCountSlice[j].count
	})

	return colorCountSlice[0:3:4]
}
