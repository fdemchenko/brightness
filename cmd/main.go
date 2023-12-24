package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	brightness "github.com/fdemchenko/brightnessctl"
)

type Settings struct {
	Value    int
	Absolute bool
	Percent  bool
	GetValue bool
}

func main() {
	settings := Settings{}
	flag.IntVar(&settings.Value, "value", 10, "brightness value")
	flag.BoolVar(&settings.Absolute, "abs", false, "set value as is instead of inc/dec by")
	flag.BoolVar(&settings.Percent, "percent", false, "use percents instead of absolute value")
	flag.BoolVar(&settings.GetValue, "get", false, "get current value")
	flag.Parse()

	bc, err := brightness.New()
	if err != nil {
		log.Fatalln(err)
	}

	if settings.GetValue {
		fmt.Println(bc.GetCurrentBrightness())
		os.Exit(0)
	}

	if settings.Percent {
		onePercent := float32(bc.GetMaxBrightness()) / 100
		if settings.Absolute {
			err = bc.SetBrightness(int(onePercent * float32(settings.Value)))
		} else {
			err = bc.SetBrightness(bc.GetCurrentBrightness() + int(onePercent*float32(settings.Value)))
		}
	} else {
		if settings.Absolute {
			err = bc.SetBrightness(settings.Value)
		} else {
			err = bc.SetBrightness(bc.GetCurrentBrightness() + settings.Value)
		}
	}
	if err != nil {
		log.Fatalln(err)
	}
}
