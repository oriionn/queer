package main

import (
	"fmt"
	"os"

	"github.com/oriionn/queer/flags"
	"golang.org/x/term"
)

type Flag uint

const (
	LesbianFlag Flag = 1 << iota
	GayFlag
	BisexualFlag
	TransgenderFlag
	PansexualFlag
	AsexualFlag
	NonBinaryFlag
	AromanticFlag
	PolysexualFlag
	DemiBoyFlag
	DemiGirlFlag
	AgenderFlag
	BigenderFlag
	GenderfluidFlag
	NeutroisFlag
	TrigenderFlag
	FemboyFlag
	ManLovingMenFlag
)

func main() {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting terminal size: %v\n", err)
		return
	}
	// fmt.Printf("Terminal size: %d columns x %d rows\n", width, height)

	flags.Lesbian(width, height)
	fmt.Println(" ")
	flags.Gay(width, height)
	fmt.Println(" ")
	flags.Bisexual(width, height)
	fmt.Println(" ")
	flags.Transgender(width, height)
}
