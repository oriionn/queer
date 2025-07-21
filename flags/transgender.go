package flags

import (
	"github.com/fatih/color"
	"github.com/oriionn/queer/flags/models"
)

func Transgender(width int, height int) {
	firstColor := color.BgRGB(91, 206, 250)
	secondColor := color.BgRGB(245, 169, 184)
	thirdColor := color.BgRGB(255, 255, 255)

	models.Pentaflag(height, width, firstColor, secondColor, thirdColor, secondColor, firstColor)
}
