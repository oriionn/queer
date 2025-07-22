package flags

import (
	"github.com/fatih/color"
	"github.com/oriionn/queer/flags/models"
)

func Genderfluid(width int, height int) {
	topColor := color.BgRGB(255, 117, 162)
	middleTopColor := color.BgRGB(255, 255, 255)
	middleColor := color.BgRGB(190, 24, 214)
	middleBottomColor := color.BgRGB(0, 0, 0)
	bottomColor := color.BgRGB(51, 62, 189)

	models.Pentaflag(height, width, topColor, middleTopColor, middleColor, middleBottomColor, bottomColor)
}
