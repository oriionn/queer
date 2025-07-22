package flags

import "github.com/fatih/color"

func Trigender() []*color.Color {
	firstColor := color.BgRGB(255, 149, 196)
	secondColor := color.BgRGB(149, 129, 254)
	thirdColor := color.BgRGB(103, 218, 97)

	return []*color.Color{firstColor, secondColor, thirdColor, secondColor, firstColor}
}
