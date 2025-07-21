package flags

import (
	"github.com/fatih/color"
	"github.com/oriionn/queer/utils"
)

func Gay(width int, height int) {
	parts := utils.Divide(6, height)

	firstColor := color.BgRGB(230, 0, 0)
	secondColor := color.BgRGB(255, 142, 0)
	thirdColor := color.BgRGB(255, 239, 0)
	fourthColor := color.BgRGB(0, 130, 27)
	fifthColor := color.BgRGB(0, 75, 255)
	sixthColor := color.BgRGB(120, 0, 137)

	colors := []*color.Color{firstColor, secondColor, thirdColor, fourthColor, fifthColor, sixthColor}
	utils.PrintFlag(width, parts, colors)
}
