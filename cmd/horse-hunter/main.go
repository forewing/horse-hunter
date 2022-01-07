package main

import (
	"flag"
	"fmt"

	hunter "github.com/forewing/horse-hunter"
)

var (
	flagLevel    = flag.String("level", hunter.LevelName[hunter.LevelDefault], "insult level, [max | min | mix]")
	flagTarget   = flag.String("target", hunter.TargetName[hunter.TargetDefault], "insult target, [female | male | mix]")
	flagInterval = flag.Duration("interval", hunter.IntervalDefault, "clipboard update interval")
)

func main() {
	flag.Parse()
	l, ok1 := hunter.LevelLookup[*flagLevel]
	if !ok1 {
		panic("invalid flag `level`")
	}

	t, ok2 := hunter.TargetLookup[*flagTarget]
	if !ok2 {
		panic("invalid flag `target`")
	}

	w := hunter.NewWorker(l, t, *flagInterval)
	if err := w.Start(); err != nil {
		panic(err)
	}
	fmt.Printf("Worker started, level: %v, target: %v, interval: %v\n",
		hunter.LevelName[w.Level], hunter.TargetName[w.Target], w.Interval)

	w.StopWaitSignal()
}
