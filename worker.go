package hunter

import (
	"errors"
	"time"

	"golang.design/x/clipboard"
)

// Worker continuously write to clipboard
type Worker struct {
	// Level of insults
	Level Level
	// Target of insults
	Target Target
	// Interval for clipboard update
	Interval time.Duration

	ch chan bool
}

const (
	WorkerIntervalMin = time.Millisecond * 10
)

var (
	ErrWorkerIntervalInvalid = errors.New("worker interval to small (<10ms)")
	ErrWorkerAlreadyStarted  = errors.New("worker already started")
	ErrWorkerAlreadyStopped  = errors.New("worker already stopped")
)

// NewWorker return a new worker
func NewWorker(level Level, target Target, interval time.Duration) *Worker {
	return &Worker{
		Level:    level,
		Target:   target,
		Interval: interval,
	}
}

// NewWorker return a new worker with default config
func NewWorkerDefault() *Worker {
	return &Worker{
		Level:    LevelDefault,
		Target:   TargetDefault,
		Interval: IntervalDefault,
	}
}

// GetLine return an insult
func (w *Worker) GetLine() string {
	return GetLine(w.Level, w.Target)
}

func ValidateWorkerInterval(interval time.Duration) error {
	if interval < WorkerIntervalMin {
		return ErrWorkerIntervalInvalid
	}
	return nil
}

// Start the worker
func (w *Worker) Start() error {
	if w.ch != nil {
		return ErrWorkerAlreadyStarted
	}

	if err := ValidateWorkerInterval(w.Interval); err != nil {
		return err
	}

	w.ch = make(chan bool)
	go w.work()

	return nil
}

// Stop the worker
func (w *Worker) Stop() error {
	if w.ch == nil {
		return ErrWorkerAlreadyStopped
	}

	close(w.ch)
	Cleanup()
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
			Cleanup()
			return
		}
	}
}

// WriteOnce to clipboard
func (w *Worker) WriteOnce() {
	clipboard.Write(clipboard.FmtText, []byte(w.GetLine()))
}

// Cleanup the clipboard
func Cleanup() {
	clipboard.Write(clipboard.FmtText, []byte(""))
}
