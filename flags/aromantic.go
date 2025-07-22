package flags

import (
	"github.com/fatih/color"
	"github.com/oriionn/queer/flags/models"
)

func Aromantic(width int, height int) {
	topColor := color.BgRGB(58, 167, 64)
	middleTopColor := color.BgRGB(168, 212, 122)
	middleColor := color.BgRGB(255, 255, 255)
	middleBottomColor := color.BgRGB(171, 171, 171)
	bottomColor := color.BgRGB(0, 0, 0)

	models.Pentaflag(height, width, topColor, middleTopColor, middleColor, middleBottomColor, bottomColor)
}
