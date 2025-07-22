package main

import (
	"fmt"
	"os"

	"github.com/oriionn/queer/colors"
	"github.com/oriionn/queer/utils"
	"golang.org/x/term"
)

type Flag uint

const (
	LesbianFlag     Flag = 1 << iota // Done
	GayFlag                          // Done
	BisexualFlag                     // Done
	TransgenderFlag                  // Done
	PansexualFlag                    // Done
	AsexualFlag                      // Done
	NonBinaryFlag                    // Done
	AromanticFlag                    // Done
	PolysexualFlag                   // Done
	DemiBoyFlag                      // Done
	DemiGirlFlag                     // Done
	AgenderFlag                      // Done
	BigenderFlag                     // Done
	GenderfluidFlag                  // Done
	NeutroisFlag                     // Done
	TrigenderFlag                    // Done
	FemboyFlag                       // Done
	MenLovingMenFlag
)

func main() {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting terminal size: %v\n", err)
		return
	}
	// fmt.Printf("Terminal size: %d columns x %d rows\n", width, height)

	utils.PrintFlag(width, height, colors.Lesbian())
	fmt.Println(" ")
	utils.PrintFlag(width, height, colors.Gay())
	fmt.Println(" ")
	utils.PrintFlag(width, height, colors.Bisexual())
	fmt.Println(" ")
	utils.PrintFlag(width, height, colors.Transgender())
	fmt.Println(" ")
	utils.PrintFlag(width, height, colors.Pansexual())
	fmt.Println(" ")
	utils.PrintFlag(width, height, colors.Asexual())
	fmt.Println(" ")
	utils.PrintFlag(width, height, colors.NonBinary())
	fmt.Println(" ")
	utils.PrintFlag(width, height, colors.Aromantic())
	fmt.Println(" ")
	utils.PrintFlag(width, height, colors.Polysexual())
	fmt.Println(" ")
	utils.PrintFlag(width, height, colors.Demiboy())
	fmt.Println(" ")
	utils.PrintFlag(width, height, colors.Demigirl())
	fmt.Println(" ")
	utils.PrintFlag(width, height, colors.Agender())
	fmt.Println(" ")
	utils.PrintFlag(width, height, colors.Bigender())
	fmt.Println(" ")
	utils.PrintFlag(width, height, colors.Genderfluid())
	fmt.Println(" ")
	utils.PrintFlag(width, height, colors.Neutrois())
	fmt.Println(" ")
	utils.PrintFlag(width, height, colors.Trigender())
	fmt.Println(" ")
	utils.PrintFlag(width, height, colors.Femboy())
	fmt.Println(" ")
	utils.PrintFlag(width, height, colors.MenLovingMen())
}
