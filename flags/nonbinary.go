package flags

import (
	"github.com/fatih/color"
	"github.com/oriionn/queer/flags/models"
)

func NonBinary(width int, height int) {
	topColor := color.BgRGB(255, 244, 47)
	secondTopColor := color.BgRGB(255, 255, 255)
	secondBottomColor := color.BgRGB(156, 89, 209)
	bottomColor := color.BgRGB(41, 41, 41)

	models.Quadriflag(height, width, topColor, secondTopColor, secondBottomColor, bottomColor)
}
