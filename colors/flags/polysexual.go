package flags

import "github.com/fatih/color"

func Polysexual() []*color.Color {
	topColor := color.BgRGB(246, 22, 186)
	middleColor := color.BgRGB(0, 214, 105)
	bottomColor := color.BgRGB(21, 147, 246)
	
	return []*color.Color{topColor, middleColor, bottomColor}
}
