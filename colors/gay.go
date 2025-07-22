package colors

import "github.com/fatih/color"

func Gay() []*color.Color {
	firstColor := color.BgRGB(230, 0, 0)
	secondColor := color.BgRGB(255, 142, 0)
	thirdColor := color.BgRGB(255, 239, 0)
	fourthColor := color.BgRGB(0, 130, 27)
	fifthColor := color.BgRGB(0, 75, 255)
	sixthColor := color.BgRGB(120, 0, 137)

	return []*color.Color{firstColor, secondColor, thirdColor, fourthColor, fifthColor, sixthColor}
}
