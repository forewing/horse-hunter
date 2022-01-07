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

```
$ horse-hunter -h
Usage of ./horse-hunter:
  -interval duration
        clipboard update interval (default 200ms)
  -level string
        insult level, [max | min | mix] (default "max")
  -target string
        insult target, [female | male | mix] (default "female")
```
