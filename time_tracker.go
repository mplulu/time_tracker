package time_tracker

import (
	"fmt"
	"math/rand"
	"time"
)

type Tracker struct {
	entries   []*TrackEntry
	startDate time.Time
	code      string
	name      string
}

func NewTracker(name string) *Tracker {
	return &Tracker{
		entries:   []*TrackEntry{},
		startDate: time.Now(),
		code:      RandSeq(5),
		name:      name,
	}
}

type TrackEntry struct {
	code      string
	startDate time.Time
}

func (t *Tracker) Track(format string, a ...interface{}) {
	code := fmt.Sprintf(format, a...)
	entry := &TrackEntry{
		code:      code,
		startDate: time.Now(),
	}
	t.entries = append(t.entries, entry)
}

func (t *Tracker) Code() string {
	return t.code
}

func (t *Tracker) OutputIfTooLong(duration time.Duration) {
	sinceStart := time.Since(t.startDate)
	if sinceStart > duration {
		t.output()
		fmt.Printf("track %s %s %s %s\n", t.name, t.code, "finished", sinceStart.String())
	}
}

func (t *Tracker) output() {
	lastDate := t.startDate
	fmt.Printf("track %s %s startAt %s\n", t.name, t.code, t.startDate)
	for _, entry := range t.entries {
		fmt.Printf("track %s %s %s %s\n", t.name, t.code, entry.code, entry.startDate.Sub(lastDate).String())
		lastDate = entry.startDate
	}

}

var letters = []rune("abcdefghijklmnopqrstuvwxyz1234567890")

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
