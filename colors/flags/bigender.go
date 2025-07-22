package flags

import "github.com/fatih/color"

func Bigender() []*color.Color {
	firstColor := color.BgRGB(196, 121, 160)
	secondColor := color.BgRGB(236, 166, 203)
	thirdColor := color.BgRGB(213, 199, 232)
	fourthColor := color.BgRGB(255, 255, 255)

	return []*color.Color{firstColor, secondColor, thirdColor, fourthColor, thirdColor, secondColor, firstColor}
}
