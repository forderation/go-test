package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
	"time"
)

const finalWord = "Go!"
const countDownStart = 3

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct{}

type SpyCountdownOperations struct {
	Calls []string
}

type ConfigurableSleeper struct {
	Duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	DurationSlept time.Duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.Duration)
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.DurationSlept = duration
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const (
	sleep = "sleep"
	write = "write"
)

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := SpySleeper{}
	Countdown(buffer, &spySleeper)
	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	Countdown(os.Stdout, &sleeper)
	if spyTime.DurationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.DurationSlept)
	}
}
