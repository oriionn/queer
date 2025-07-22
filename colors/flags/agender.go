package flags

import "github.com/fatih/color"

func Agender() []*color.Color {
	firstColor := color.BgRGB(0, 0, 0)
	secondColor := color.BgRGB(188, 196, 198)
	thirdColor := color.BgRGB(255, 255, 255)
	fourthColor := color.BgRGB(182, 245, 131)

	return []*color.Color{firstColor, secondColor, thirdColor, fourthColor, thirdColor, secondColor, firstColor}
}
