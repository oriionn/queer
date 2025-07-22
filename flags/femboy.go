package flags

import (
	"github.com/fatih/color"
	"github.com/oriionn/queer/flags/models"
)

func Femboy(width int, height int) {
	firstColor := color.BgRGB(213, 96, 169)
	secondColor := color.BgRGB(229, 175, 206)
	thirdColor := color.BgRGB(255, 255, 255)
	fourthColor := color.BgRGB(86, 207, 248)

	models.Heptaflag(height, width, firstColor, secondColor, thirdColor, fourthColor, thirdColor, secondColor, firstColor)
}
