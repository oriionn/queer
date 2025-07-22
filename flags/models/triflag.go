package models

import (
	"github.com/fatih/color"
	"github.com/oriionn/queer/utils"
)

func Triflag(height int, width int, topColor *color.Color, middleColor *color.Color, bottomColor *color.Color) {
	parts := utils.Divide(3, height)
	colors := []*color.Color{topColor, middleColor, bottomColor}
	utils.PrintFlag(width, parts, colors)
}
