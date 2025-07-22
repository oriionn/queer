package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintFlag(width int, height int, colors []*color.Color) {
	parts := Divide(len(colors), height)

	for _, color := range colors {
		for range parts {
			fmt.Println(color.Sprintf("%s", Repeat(" ", width)))
		}
	}
}
