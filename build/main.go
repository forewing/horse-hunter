package main

import (
	"flag"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/forewing/gobuild"
)

const (
	outputPath = "./output"
	appName    = "HorseHunter"

	cliSource = "./cmd/horse-hunter"
	cliOutput = "horse-hunter"

	guiSource = "./cmd/horse-hunter-gui"
	guiOutput = "HorseHunter-GUI"
)

var (
	flagZip = flag.Bool("zip", false, "compress output")
	flagCLI = flag.Bool("cli", false, "build CLI")
	flagGUI = flag.Bool("gui", false, "build GUI")

	target = gobuild.Target{
		Source:     cliSource,
		OutputName: cliOutput,
		OutputPath: outputPath,

		CleanOutput: true,
		Cgo:         true,

		ExtraFlags:   []string{"-trimpath"},
		ExtraLdFlags: "-s -w",

		Compress: gobuild.CompressRaw,
		Platforms: []gobuild.Platform{{
			OS:   gobuild.PlatformOS(runtime.GOOS),
			Arch: gobuild.PlatformArch(runtime.GOARCH),
		}},
	}
)

func main() {
	flag.Parse()

	if *flagZip {
		target.Compress = gobuild.CompressAllZip
	}

	if !*flagCLI && !*flagGUI {
		*flagCLI = true
		*flagGUI = true
	}

	if runtime.GOOS == "darwin" {
		target.Platforms[0].Arch = gobuild.ArchUniversal
	}

	if *flagCLI {
		err := target.Build()
		if err != nil {
			panic(err)
		}
	}

	if *flagGUI {
		switch runtime.GOOS {
		case "windows":
			target.ExtraLdFlags += " -H=windowsgui"
		case "darwin":
			target.Compress = gobuild.CompressRaw
		}

		if *flagCLI {
			target.CleanOutput = false
		}

		target.Source = guiSource
		target.OutputName = guiOutput
		err := target.Build()
		if err != nil {
			panic(err)
		}

		if runtime.GOOS == "darwin" {
			err := gobuild.BuildMacOSApp(
				outputPath,
				appName,
				filepath.Join(outputPath, guiOutput),
				"com.github.forewing.HorseHunter",
				filepath.Join(guiSource, "resources/icon.png"),
				true,
			)
			if err != nil {
				panic(err)
			}

			if *flagZip {
				os.Chdir(outputPath)
				exec.Command("zip", "-qmr", "HorseHunter-GUI-macOS-universal.zip", appName+".app").Run()
			}
		}
	}
}
