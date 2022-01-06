package hunter

import (
	"errors"
	"time"

	"golang.design/x/clipboard"
)

type Worker struct {
	ch       chan bool
	Level    Level
	Target   Target
	Interval time.Duration
}

var (
	ErrWorkerAlreadyStarted = errors.New("worker already started")
	ErrWorkerAlreadyStopped = errors.New("worker already stopped")
)

func NewWorker(level Level, target Target, interval time.Duration) *Worker {
	return &Worker{
		Level:    level,
		Target:   target,
		Interval: interval,
	}
}

func (w *Worker) GetLine() string {
	return GetLine(w.Level, w.Target)
}

func (w *Worker) Start() error {
	if w.ch != nil {
		return ErrWorkerAlreadyStarted
	}

	w.ch = make(chan bool)
	go w.work()

	return nil
}

func (w *Worker) Stop() error {
	if w.ch == nil {
		return ErrWorkerAlreadyStopped
	}

	close(w.ch)
	w.Cleanup()
	return nil
}

func (w *Worker) work() {
	ticker := time.NewTicker(w.Interval)
	w.WriteOnce()
	for {
		select {
		case <-ticker.C:
			w.WriteOnce()
		case <-w.ch:
			ticker.Stop()
			w.ch = nil
			w.Cleanup()
			return
		}
	}
}

func (w *Worker) WriteOnce() {
	clipboard.Write(clipboard.FmtText, []byte(w.GetLine()))
}

func (w *Worker) Cleanup() {
	clipboard.Write(clipboard.FmtText, []byte(""))
}
