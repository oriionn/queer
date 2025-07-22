package main

import (
	"fmt"
	"os"

	"github.com/oriionn/queer/flags"
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

	flags.Lesbian(width, height)
	fmt.Println(" ")
	flags.Gay(width, height)
	fmt.Println(" ")
	flags.Bisexual(width, height)
	fmt.Println(" ")
	flags.Transgender(width, height)
	fmt.Println(" ")
	flags.Pansexual(width, height)
	fmt.Println(" ")
	flags.Asexual(width, height)
	fmt.Println(" ")
	flags.NonBinary(width, height)
	fmt.Println(" ")
	flags.Aromantic(width, height)
	fmt.Println(" ")
	flags.Polysexual(width, height)
	fmt.Println(" ")
	flags.Demiboy(width, height)
	fmt.Println(" ")
	flags.Demigirl(width, height)
	fmt.Println(" ")
	flags.Agender(width, height)
	fmt.Println(" ")
	flags.Bigender(width, height)
	fmt.Println(" ")
	flags.Genderfluid(width, height)
	fmt.Println(" ")
	flags.Neutrois(width, height)
	fmt.Println(" ")
	flags.Trigender(width, height)
	fmt.Println(" ")
	flags.Femboy(width, height)
	fmt.Println(" ")
	flags.MenLovingMen(width, height)
}
