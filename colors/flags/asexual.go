package flags

import "github.com/fatih/color"

func Asexual() []*color.Color {
	topColor := color.BgRGB(0, 0, 0)
	secondTopColor := color.BgRGB(165, 165, 165)
	secondBottomColor := color.BgRGB(255, 255, 255)
	bottomColor := color.BgRGB(129, 0, 129)

	return []*color.Color{topColor, secondTopColor, secondBottomColor, bottomColor}
}
