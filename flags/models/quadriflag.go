package models

import (
	"github.com/fatih/color"
	"github.com/oriionn/queer/utils"
)

func Quadriflag(height int, width int, topColor *color.Color, secondTopColor *color.Color, secondBottomColor *color.Color, bottomColor *color.Color) {
	parts := utils.Divide(4, height)
	colors := []*color.Color{topColor, secondTopColor, secondBottomColor, bottomColor}
	utils.PrintFlag(width, parts, colors)
}
