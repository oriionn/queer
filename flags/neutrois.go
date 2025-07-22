package flags

import (
	"github.com/fatih/color"
	"github.com/oriionn/queer/flags/models"
)

func Neutrois(width int, height int) {
	topColor := color.BgRGB(255, 255, 255)
	middleColor := color.BgRGB(54, 159, 36)
	bottomColor := color.BgRGB(8, 8, 8)

	models.Triflag(height, width, topColor, middleColor, bottomColor)
}
