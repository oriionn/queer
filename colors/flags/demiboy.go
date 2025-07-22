package flags

import "github.com/fatih/color"

func Demiboy() []*color.Color {
	firstColor := color.BgRGB(128, 128, 128)
	secondColor := color.BgRGB(197, 197, 197)
	thirdColor := color.BgRGB(155, 218, 236)
	fourthColor := color.BgRGB(255, 255, 255)

	return []*color.Color{firstColor, secondColor, thirdColor, fourthColor, thirdColor, secondColor, firstColor}
}
