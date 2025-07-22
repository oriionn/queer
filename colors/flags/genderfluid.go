package flags

import "github.com/fatih/color"

func Genderfluid() []*color.Color {
	topColor := color.BgRGB(255, 117, 162)
	middleTopColor := color.BgRGB(255, 255, 255)
	middleColor := color.BgRGB(190, 24, 214)
	middleBottomColor := color.BgRGB(0, 0, 0)
	bottomColor := color.BgRGB(51, 62, 189)

	return []*color.Color{topColor, middleTopColor, middleColor, middleBottomColor, bottomColor}
}
