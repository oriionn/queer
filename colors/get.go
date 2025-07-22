package colors

import (
	"github.com/fatih/color"
	"github.com/oriionn/queer/colors/flags"
)

func Get(c Color) []*color.Color {
	switch c {
	case Agender:
		return flags.Agender()
	case Aromantic:
		return flags.Aromantic()
	case Asexual:
		return flags.Asexual()
	case Bigender:
		return flags.Bigender()
	case Bisexual:
		return flags.Bisexual()
	case DemiBoy:
		return flags.Demiboy()
	case DemiGirl:
		return flags.Demigirl()
	case Femboy:
		return flags.Femboy()
	case Gay:
		return flags.Gay()
	case Genderfluid:
		return flags.Genderfluid()
	case Lesbian:
		return flags.Lesbian()
	case MenLovingMen:
		return flags.MenLovingMen()
	case Neutrois:
		return flags.Neutrois()
	case NonBinary:
		return flags.NonBinary()
	case Pansexual:
		return flags.Pansexual()
	case Polysexual:
		return flags.Polysexual()
	case Transgender:
		return flags.Transgender()
	case Trigender:
		return flags.Trigender()
	default:
		return flags.Gay()
	}
}
