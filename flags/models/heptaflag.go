package models

import (
	"github.com/fatih/color"
	"github.com/oriionn/queer/utils"
)

func Heptaflag(height int, width int, topColor *color.Color, middleTopColor *color.Color, middleMiddleTopColor *color.Color, middleColor *color.Color, middleMiddleBottomColor *color.Color, middleBottomColor *color.Color, bottomColor *color.Color) {
	parts := utils.Divide(7, height)

	colors := []*color.Color{topColor, middleTopColor, middleMiddleTopColor, middleColor, middleMiddleBottomColor, middleBottomColor, bottomColor}
	utils.PrintFlag(width, parts, colors)
}
