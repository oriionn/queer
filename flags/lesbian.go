package flags

import (
	"github.com/fatih/color"
	"github.com/oriionn/queer/flags/models"
)

func Lesbian(width int, height int) {
	topColor := color.BgRGB(212, 43, 0)
	middleTopColor := color.BgRGB(254, 153, 85)
	middleColor := color.BgRGB(255, 255, 255)
	middleBottomColor := color.BgRGB(210, 97, 163)
	bottomColor := color.BgRGB(162, 1, 97)

	models.Pentaflag(height, width, topColor, middleTopColor, middleColor, middleBottomColor, bottomColor)
}
