package main

import (
	"flag"
	"os"
	"path/filepath"
	"runtime"

	"github.com/forewing/gobuild"
	"github.com/mholt/archiver/v4"
)

const (
	outputPath = "./output"
	appName    = "HorseHunter"
	appZipName = "HorseHunter-GUI-macOS-universal.zip"

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
		target.Compress = gobuild.CompressZip
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
				appPath := filepath.Join(outputPath, appName+".app")
				err := gobuild.Compress(
					filepath.Join(outputPath, appZipName),
					map[string]string{appPath: ""},
					archiver.Zip{})
				if err != nil {
					panic(err)
				}
				os.RemoveAll(appPath)
			}
		}
	}
}
