package ping

import (
	"errors"
	"time"

	pping "github.com/sparrc/go-ping"
)

type Latency struct {
	Duration time.Duration
	Loss     float64
}

// TODO: take packet loss into account
func (l Latency) Less(other Latency) bool {
	return l.Duration < other.Duration
}

// TODO: add configuration for ping, e.g. timeout/count

// Latency returns the avg ping latency for the addr
// addr can be ip string or domain name
func GetLatency(addr string) (Latency, error) {
	pinger, err := pping.NewPinger(addr)
	if err != nil {
		return Latency{}, err
	}

	pinger.Count = 3
	pinger.Timeout = 2 * time.Second
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats

	if stats.PacketLoss == 100 {
		return Latency{}, errors.New("packets all lost")
	}

	return Latency{
		Duration: stats.AvgRtt,
		Loss:     stats.PacketLoss,
	}, nil
}
