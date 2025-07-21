package flags

import (
	"github.com/fatih/color"
	"github.com/oriionn/queer/flags/models"
)

func Bisexual(width int, height int) {
	topColor := color.BgRGB(215, 0, 113)
	middleColor := color.BgRGB(156, 78, 151)
	bottomColor := color.BgRGB(0, 53, 170)

	models.Pentaflag(height, width, topColor, topColor, middleColor, bottomColor, bottomColor)
}
