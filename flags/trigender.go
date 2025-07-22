package flags

import (
	"github.com/fatih/color"
	"github.com/oriionn/queer/flags/models"
)

func Trigender(width int, height int) {
	firstColor := color.BgRGB(255, 149, 196)
	secondColor := color.BgRGB(149, 129, 254)
	thirdColor := color.BgRGB(103, 218, 97)

	models.Pentaflag(height, width, firstColor, secondColor, thirdColor, secondColor, firstColor)
}
