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
