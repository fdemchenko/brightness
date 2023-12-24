package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

type Settings struct {
	Value    int
	Absolute bool
	Percent  bool
	GetValue bool
}

const (
	backlightDir          = "/sys/class/backlight/intel_backlight"
	brightnessFilename    = "brightness"
	maxBrightnessFilename = "max_brightness"
)

func main() {
	settings := Settings{}
	flag.IntVar(&settings.Value, "value", 10, "brightness value")
	flag.BoolVar(&settings.Absolute, "abs", false, "set value as is, or inc/dec by")
	flag.BoolVar(&settings.Percent, "percent", false, "use percents instead of absolute value")
	flag.BoolVar(&settings.GetValue, "get", false, "use percents instead of absolute value")
	flag.Parse()

	brightness, err := getCurrentBrightness()
	if err != nil {
		log.Fatalln(err)
	}

	if settings.GetValue {
		fmt.Println(brightness)
		os.Exit(1)
	}

	maxBrightness, err := getMaxBrightness()
	if err != nil {
		log.Fatalln(err)
	}

	if settings.Percent {
		onePercent := float32(maxBrightness) / 100
		if settings.Absolute {
			brightness = int(onePercent * float32(settings.Value))
		} else {
			brightness += int(onePercent * float32(settings.Value))
		}
	} else {
		if settings.Absolute {
			brightness = settings.Value
		} else {
			brightness += settings.Value
		}
	}

	brightness = checkBounds(brightness, 0, maxBrightness)
	err = os.WriteFile(path.Join(backlightDir, brightnessFilename), []byte(fmt.Sprintf("%d\n", brightness)), 0660)
	if err != nil {
		log.Fatalln(err)
	}
}

func getMaxBrightness() (int, error) {
	filename := path.Join(backlightDir, maxBrightnessFilename)
	return getNumberFromFile(filename)
}

func getCurrentBrightness() (int, error) {
	filename := path.Join(backlightDir, brightnessFilename)
	return getNumberFromFile(filename)
}
