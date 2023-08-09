package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/atomic"
	"math"
	"time"
)

type TimeoutGauge struct {
	G              prometheus.Gauge
	expiryDuration time.Duration
	lastUpdate     *atomic.Time
}

func NewTimeoutGauge(expiryDuration time.Duration, gaugeOpts prometheus.GaugeOpts) *TimeoutGauge {
	g := prometheus.NewGauge(gaugeOpts)
	g.Set(math.NaN())
	t := &TimeoutGauge{
		G:              g,
		expiryDuration: expiryDuration,
		lastUpdate:     atomic.NewTime(time.Now()),
	}
	go t.pollTimeout()

	return t
}

func (t *TimeoutGauge) Set(val float64) {
	t.G.Set(val)
	t.lastUpdate.Store(time.Now())
}

func (t *TimeoutGauge) pollTimeout() {
	for tick := range time.NewTicker(time.Second).C {
		if tick.Sub(t.lastUpdate.Load()) > t.expiryDuration {
			t.unset()
		}
	}
}

func (t *TimeoutGauge) unset() {
	t.G.Set(math.NaN())
}
