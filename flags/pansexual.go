package flags

import (
	"github.com/fatih/color"
	"github.com/oriionn/queer/flags/models"
)

func Pansexual(width int, height int) {
	topColor := color.BgRGB(255, 27, 141)
	middleColor := color.BgRGB(255, 217, 0)
	bottomColor := color.BgRGB(27, 179, 255)

	models.Triflag(height, width, topColor, middleColor, bottomColor)
}
