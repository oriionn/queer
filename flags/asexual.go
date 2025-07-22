package flags

import (
	"github.com/fatih/color"
	"github.com/oriionn/queer/flags/models"
)

func Asexual(width int, height int) {
	topColor := color.BgRGB(0, 0, 0)
	secondTopColor := color.BgRGB(165, 165, 165)
	secondBottomColor := color.BgRGB(255, 255, 255)
	bottomColor := color.BgRGB(129, 0, 129)

	models.Quadriflag(height, width, topColor, secondTopColor, secondBottomColor, bottomColor)
}
