package flags

import "github.com/fatih/color"

func Demigirl() []*color.Color {
	firstColor := color.BgRGB(127, 127, 127)
	secondColor := color.BgRGB(195, 195, 195)
	thirdColor := color.BgRGB(255, 174, 201)
	fourthColor := color.BgRGB(255, 255, 255)

	return []*color.Color{firstColor, secondColor, thirdColor, fourthColor, thirdColor, secondColor, firstColor}
}
