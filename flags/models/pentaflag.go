package models

import (
	"github.com/fatih/color"
	"github.com/oriionn/queer/utils"
)

func Pentaflag(height int, width int, topColor *color.Color, middleTopColor *color.Color, middleColor *color.Color, middleBottomColor *color.Color, bottomColor *color.Color) {
	parts := utils.Divide(5, height)

	colors := []*color.Color{topColor, middleTopColor, middleColor, middleBottomColor, bottomColor}
	utils.PrintFlag(width, parts, colors)
}
