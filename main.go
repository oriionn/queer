package main

import (
	"flag"
	"fmt"
	"math/rand"
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
	{[]string{"agender", "ag"}, colors.Agender, "print the agender flag"},
	{[]string{"aromantic", "ar"}, colors.Aromantic, "print the aromantic flag"},
	{[]string{"asexual", "as"}, colors.Asexual, "print the asexual flag"},
	{[]string{"bigender", "big"}, colors.Bigender, "print the bigender flag"},
	{[]string{"bisexual", "bi", "b"}, colors.Bisexual, "print the bisexual flag"},
	{[]string{"demiboy", "db"}, colors.DemiBoy, "print the demiboy flag"},
	{[]string{"demigirl", "dg"}, colors.DemiGirl, "print the demigirl flag"},
	{[]string{"femboy", "f"}, colors.Femboy, "print the femboy flag"},
	{[]string{"gay", "g"}, colors.Gay, "print the gay flag"},
	{[]string{"genderfluid", "ge"}, colors.Genderfluid, "print the genderfluid flag"},
	{[]string{"lesbian", "l"}, colors.Lesbian, "print the lesbian flag"},
	{[]string{"menlovingmen", "mlm"}, colors.MenLovingMen, "print the men loving men flag"},
	{[]string{"neutrois", "ne"}, colors.Neutrois, "print the neutrois flag"},
	{[]string{"nonbinary", "nb"}, colors.NonBinary, "print the non binary flag"},
	{[]string{"pansexual", "pan", "pa"}, colors.Pansexual, "print the pansexual flag"},
	{[]string{"polysexual", "poly", "po"}, colors.Polysexual, "print the polysexual flag"},
	{[]string{"transgender", "trans", "t"}, colors.Transgender, "print the transgender flag"},
	{[]string{"trigender", "tr"}, colors.Trigender, "print the trigender flag"},
}

func main() {
	var (
		color colors.Color = colors.Undefined
		help  bool
	)

	for _, arg := range args {
		for _, name := range arg.args {
			flag.BoolFunc(name, "", func(_ string) error {
				color = arg.value
				return nil
			})
		}
	}

	flag.BoolVar(&help, "help", false, "")
	flag.BoolVar(&help, "h", false, "")

	flag.Parse()

	if color == colors.Undefined {
		index := rand.Intn(len(args))
		color = args[index].value
	}

	if help == true {
		printHelp()
		return
	}

	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting terminal size: %v\n", err)
		return
	}

	utils.PrintFlag(width, height, colors.Get(color))
}

func printHelp() {
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  queer [options]")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  -h, --help                     show this help menu")
	fmt.Println("")
	for _, arg := range args {
		fmt.Printf("")

		argsList := ""
		for i, name := range arg.args {
			if i == 0 {
				argsList += fmt.Sprintf("--%s", name)
			} else {
				argsList += fmt.Sprintf(", -%s", name)
			}
		}
		fmt.Printf("%s", argsList)

		fmt.Printf("%s %s\n", utils.Repeat(" ", 30-len(argsList)), arg.description)
	}
	fmt.Println("")
}
