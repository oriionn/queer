package colors

import "github.com/fatih/color"

func Pansexual() []*color.Color {
	topColor := color.BgRGB(255, 27, 141)
	middleColor := color.BgRGB(255, 217, 0)
	bottomColor := color.BgRGB(27, 179, 255)

	return []*color.Color{topColor, middleColor, bottomColor}
}
