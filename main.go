package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/oriionn/queer/colors"
	"github.com/oriionn/queer/utils"
	"golang.org/x/term"
)

type Argument struct {
	args        []string
	value       colors.Color
	description string
}

var args []Argument = []Argument{
	{[]string{"agender", "ag"}, colors.Agender, "Print the agender flag"},
	{[]string{"aromantic", "ar"}, colors.Aromantic, "Print the aromantic flag"},
	{[]string{"asexual", "as"}, colors.Asexual, "Print the asexual flag"},
	{[]string{"bigender", "big"}, colors.Bigender, "Print the bigender flag"},
	{[]string{"bisexual", "bi", "b"}, colors.Bisexual, "Print the bisexual flag"},
	{[]string{"demiboy", "db"}, colors.DemiBoy, "Print the demiboy flag"},
	{[]string{"demigirl", "dg"}, colors.DemiGirl, "Print the demigirl flag"},
	{[]string{"femboy", "f"}, colors.Femboy, "Print the femboy flag"},
	{[]string{"gay", "g"}, colors.Gay, "Print the gay flag"},
	{[]string{"genderfluid", "ge"}, colors.Genderfluid, "Print the genderfluid flag"},
	{[]string{"lesbian", "l"}, colors.Lesbian, "Print the lesbian flag"},
	{[]string{"menlovingmen", "mlm"}, colors.MenLovingMen, "Print the men loving men flag"},
	{[]string{"neutrois", "ne"}, colors.Neutrois, "Print the neutrois flag"},
	{[]string{"nonbinary", "nb"}, colors.NonBinary, "Print the non binary flag"},
	{[]string{"pansexual", "pan", "pa"}, colors.Pansexual, "Print the pansexual flag"},
	{[]string{"polysexual", "poly", "po"}, colors.Polysexual, "Print the polysexual flag"},
	{[]string{"transgender", "trans", "t"}, colors.Transgender, "Print the transgender flag"},
	{[]string{"trigender", "tr"}, colors.Trigender, "Print the trigender flag"},
}

func main() {
	var color colors.Color

	for _, arg := range args {
		for _, name := range arg.args {
			flag.BoolFunc(name, "", func(_ string) error {
				color = arg.value
				return nil
			})
		}
	}

	flag.Parse()

	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting terminal size: %v\n", err)
		return
	}

	utils.PrintFlag(width, height, colors.Get(color))
}
