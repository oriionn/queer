package colors

import "github.com/fatih/color"

func Femboy() []*color.Color {
	firstColor := color.BgRGB(213, 96, 169)
	secondColor := color.BgRGB(229, 175, 206)
	thirdColor := color.BgRGB(255, 255, 255)
	fourthColor := color.BgRGB(86, 207, 248)

	return []*color.Color{firstColor, secondColor, thirdColor, fourthColor, thirdColor, secondColor, firstColor}
}
