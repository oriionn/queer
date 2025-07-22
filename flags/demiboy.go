package flags

import (
	"github.com/fatih/color"
	"github.com/oriionn/queer/flags/models"
)

func Demiboy(width int, height int) {
	firstColor := color.BgRGB(128, 128, 128)
	secondColor := color.BgRGB(197, 197, 197)
	thirdColor := color.BgRGB(155, 218, 236)
	fourthColor := color.BgRGB(255, 255, 255)

	models.Heptaflag(height, width, firstColor, secondColor, thirdColor, fourthColor, thirdColor, secondColor, firstColor)
}
