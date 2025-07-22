package flags

import (
	"github.com/fatih/color"
	"github.com/oriionn/queer/flags/models"
)

func Polysexual(width int, height int) {
	topColor := color.BgRGB(246, 22, 186)
	middleColor := color.BgRGB(0, 214, 105)
	bottomColor := color.BgRGB(21, 147, 246)

	models.Triflag(height, width, topColor, middleColor, bottomColor)
}
