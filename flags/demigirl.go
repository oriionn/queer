package flags

import (
	"github.com/fatih/color"
	"github.com/oriionn/queer/flags/models"
)

func Demigirl(width int, height int) {
	firstColor := color.BgRGB(127, 127, 127)
	secondColor := color.BgRGB(195, 195, 195)
	thirdColor := color.BgRGB(255, 174, 201)
	fourthColor := color.BgRGB(255, 255, 255)

	models.Heptaflag(height, width, firstColor, secondColor, thirdColor, fourthColor, thirdColor, secondColor, firstColor)
}
