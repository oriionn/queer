package colors

import "github.com/fatih/color"

func NonBinary() []*color.Color {
	topColor := color.BgRGB(255, 244, 47)
	secondTopColor := color.BgRGB(255, 255, 255)
	secondBottomColor := color.BgRGB(156, 89, 209)
	bottomColor := color.BgRGB(41, 41, 41)

	return []*color.Color{topColor, secondTopColor, secondBottomColor, bottomColor}
}
