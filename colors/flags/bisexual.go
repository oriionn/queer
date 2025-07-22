package flags

import "github.com/fatih/color"

func Bisexual() []*color.Color {
	topColor := color.BgRGB(215, 0, 113)
	middleColor := color.BgRGB(156, 78, 151)
	bottomColor := color.BgRGB(0, 53, 170)

	return []*color.Color{topColor, topColor, middleColor, bottomColor, bottomColor}
}
