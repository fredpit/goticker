/*
 * GO Ticker
 *
 * Frédéric PITOIZET.	2020-12-01.
 *
 */

package goticker

import "time"

// Ticker : struct containing ticker
type Ticker struct {
	nextTick     time.Time
	tickDuration time.Duration
}

// New : create a ticker with no delay at start
func New(d time.Duration) Ticker {
	return NewWithInitialDelay(d, 0)
}

// NewWithInitialDelay : create a ticler with an initial delay
func NewWithInitialDelay(d, dl time.Duration) Ticker {
	return Ticker{
		nextTick:     time.Now().Add(dl),
		tickDuration: d,
	}
}

// Tick : test if the ticker has ticked
func (t Ticker) Tick() bool {
	if time.Now().After(t.nextTick) {
		t.Reset()
		return true
	}
	return false
}

// Reset : restart tick delay
func (t Ticker) Reset() {
	t.nextTick = time.Now().Add(t.tickDuration)
}
