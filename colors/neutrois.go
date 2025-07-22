package colors

import "github.com/fatih/color"

func Neutrois() []*color.Color {
	topColor := color.BgRGB(255, 255, 255)
	middleColor := color.BgRGB(54, 159, 36)
	bottomColor := color.BgRGB(8, 8, 8)

	return []*color.Color{topColor, middleColor, bottomColor}
}
