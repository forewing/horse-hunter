package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	hunter "github.com/forewing/horse-hunter"
)

const (
	textStopped = "Stopped"
	textRunning = "Running"
)

var (
	worker = hunter.NewWorkerDefault()

	intervalInput *widget.Entry
	levelSelect   *widget.RadioGroup
	targetSelect  *widget.RadioGroup
	statusLabel   *widget.Label

	radioPadding = 0
)

func init() {
	for _, s := range append(hunter.LevelName, hunter.TargetName...) {
		if len(s) > radioPadding {
			radioPadding = len(s)
		}
	}
}

func main() {
	go func() {
		worker.StopWaitSignal()
		os.Exit(0)
	}()

	myApp := app.New()
	myWindow := myApp.NewWindow("Horse Hunter")
	myApp.Settings().SetTheme(&myTheme{})
	myWindow.SetIcon(myIcon{})

	setupWidgets()
	rows := []fyne.CanvasObject{
		widget.NewLabel("Level"),
		levelSelect,

		widget.NewLabel("Target"),
		targetSelect,

		widget.NewLabel("Interval(s)"),
		intervalInput,

		statusLabel,
		widget.NewButton("Start/Stop", toggleWorker),
	}

	grid := container.New(layout.NewFormLayout(), rows...)
	myWindow.SetContent(grid)
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()

	worker.Stop()
	hunter.Cleanup()
}

func setupWidgets() {
	intervalInput = widget.NewEntry()
	intervalInput.SetText("0.2")
	intervalInput.Validator = func(s string) error {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		fd := time.Duration((math.Floor(f * 1000))) * time.Millisecond
		if err := hunter.ValidateWorkerInterval(fd); err != nil {
			return err
		}
		worker.Interval = fd
		return nil
	}

	levelSelect = widget.NewRadioGroup(padRadioTextSlice(hunter.LevelName), func(s string) {
		worker.Level = hunter.LevelLookup[strings.TrimSpace(s)]
	})
	levelSelect.Horizontal = true
	levelSelect.Required = true
	levelSelect.SetSelected(padRadioText(hunter.LevelName[hunter.LevelDefault]))

	targetSelect = widget.NewRadioGroup(padRadioTextSlice(hunter.TargetName), func(s string) {
		worker.Target = hunter.TargetLookup[strings.TrimSpace(s)]
	})
	targetSelect.Horizontal = true
	targetSelect.Required = true
	targetSelect.SetSelected(padRadioText(hunter.TargetName[hunter.TargetDefault]))

	statusLabel = widget.NewLabel(textStopped)
}

func toggleWorker() {
	defer func() {
		if worker.IsRunning() {
			statusLabel.SetText(textRunning)
		} else {
			statusLabel.SetText(textStopped)
		}
	}()

	if worker.IsRunning() {
		log.Println("stop")
		if err := worker.Stop(); err != nil {
			log.Println(err)
		}
	} else {
		log.Printf("start: level: %v, target: %v, interval: %v", worker.Level, worker.Target, worker.Interval)
		if err := intervalInput.Validate(); err != nil {
			log.Println(err)
		} else if err := worker.Start(); err != nil {
			log.Println(err)
		}
	}
}

func padRadioTextSlice(ss []string) []string {
	for i := range ss {
		ss[i] = padRadioText(ss[i])
	}
	return ss
}

func padRadioText(s string) string {
	return s + strings.Repeat(" ", radioPadding-len(s))
}
