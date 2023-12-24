package brightness

import (
	"fmt"
	"os"
	"path"
)

const (
	backlightDir          = "/sys/class/backlight/intel_backlight"
	brightnessFilename    = "brightness"
	maxBrightnessFilename = "max_brightness"
)

type BrightnessController struct {
	currentBrightness, maxBrightness int
}

func getMaxBrightness() (int, error) {
	filename := path.Join(backlightDir, maxBrightnessFilename)
	return getNumberFromFile(filename)
}

func getCurrentBrightness() (int, error) {
	filename := path.Join(backlightDir, brightnessFilename)
	return getNumberFromFile(filename)
}

func New() (*BrightnessController, error) {
	currentBrightness, err := getCurrentBrightness()
	if err != nil {
		return nil, err
	}

	maxBrightness, err := getMaxBrightness()
	if err != nil {
		return nil, err
	}
	return &BrightnessController{maxBrightness: maxBrightness, currentBrightness: currentBrightness}, nil
}

func (bc *BrightnessController) GetCurrentBrightness() int {
	return bc.currentBrightness
}

func (bc *BrightnessController) GetMaxBrightness() int {
	return bc.maxBrightness
}

func (bc *BrightnessController) SetBrightness(value int) error {
	value = checkBounds(value, 0, bc.maxBrightness)
	err := os.WriteFile(path.Join(backlightDir, brightnessFilename), []byte(fmt.Sprintf("%d\n", value)), 0660)
	if err != nil {
		bc.currentBrightness = value
	}
	return err
}
