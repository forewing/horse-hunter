package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	hunter "github.com/forewing/go-horse-hunter"
)

var (
	flagLevel    = flag.String("level", "max", "insult level, [max | min | mix]")
	flagTarget   = flag.String("target", "female", "insult target, [female | male | mix]")
	flagInterval = flag.Duration("interval", time.Millisecond*200, "clipboard update interval")
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

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
	w.Stop()
}
