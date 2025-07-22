package colors

import "github.com/fatih/color"

func Transgender() []*color.Color {
	firstColor := color.BgRGB(91, 206, 250)
	secondColor := color.BgRGB(245, 169, 184)
	thirdColor := color.BgRGB(255, 255, 255)

	return []*color.Color{firstColor, secondColor, thirdColor, secondColor, firstColor}
}
