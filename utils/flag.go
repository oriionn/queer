package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintFlag(width int, parts int, colors []*color.Color) {
	for _, color := range colors {
		for range parts {
			fmt.Println(color.Sprintf("%s", Repeat(" ", width)))
		}
	}
}
