package flags

import "github.com/fatih/color"

func MenLovingMen() []*color.Color {
	topColor := color.BgRGB(0, 139, 109)
	middleTopColor := color.BgRGB(27, 203, 165)
	middleMiddleTopColor := color.BgRGB(153, 240, 197)
	middleColor := color.BgRGB(255, 255, 255)
	middleMiddleBottomColor := color.BgRGB(112, 163, 215)
	middleBottomColor := color.BgRGB(75, 67, 212)
	bottomColor := color.BgRGB(55, 15, 122)

	return []*color.Color{topColor, middleTopColor, middleMiddleTopColor, middleColor, middleMiddleBottomColor, middleBottomColor, bottomColor}
}
