# Horse Hunter (Go Implementation)

[![PkgGoDev](https://pkg.go.dev/badge/github.com/forewing/horse-hunter)](https://pkg.go.dev/github.com/forewing/horse-hunter)

## API

```go
// get a single insult by arguments
fmt.Println(hunter.GetLine(hunter.LevelMax, hunter.TargetMale))

// create a worker
worker := hunter.NewWorkerDefault()

// get a single insult
fmt.Println(worker.GetLine())

// write an insult to clipboard
worker.WriteOnce()

// start the worker, continuously write insult to clipboard
worker.Start()

// stop the worker
worker.Stop()

// clean up the clipboard
hunter.Cleanup()
```

## CLI

> On Linux, you need `libx11-dev` installed.

```shell
# Build
go build ./cmd/horse-hunter

# Run
$ horse-hunter -h
Usage of ./horse-hunter:
  -interval duration
        clipboard update interval (default 200ms)
  -level string
        insult level, [max | min | mix] (default "max")
  -target string
        insult target, [female | male | mix] (default "female")
```

## GUI

<img src="https://user-images.githubusercontent.com/13747187/148613206-508d9ed9-a952-4ea7-9a1d-8aeecdbb9c18.png" alt="showcase" width="50%"/>

### Simple executable

```shell
go build ./cmd/horse-hunter-gui
```

### Package for distribution

1. Windows

```
go build -trimpath -ldflags "-s -w -H=windowsgui" -o HorseHunter.exe ./cmd/horse-hunter-gui
```

2. macOS

```shell
# Install fyne cli tools
go install fyne.io/fyne/v2/cmd/fyne@latest

# Build binary
go build -trimpath -ldflags "-s -w" ./cmd/horse-hunter-gui

# Bundle package
fyne package -icon ./cmd/horse-hunter-gui/icon.png -name HorseHunter -release -exe horse-hunter-gui
```

3. Linux

```shell
# Install dependencies
sudo apt install xorg-dev -y

# Build binary
go build -trimpath -ldflags "-s -w" -o HorseHunter ./cmd/horse-hunter-gui
```
