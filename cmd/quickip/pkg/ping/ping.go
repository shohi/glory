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

const (
	defaultCount   = 3
	defaultTimeout = 2 * time.Second
)

// Options for ping
type Options struct {
	Count   int
	Timeout time.Duration
}

// DefaultOptions returns default configuration options for ping.
func DefaultOptions() Options {
	return Options{
		Count:   defaultCount,
		Timeout: defaultTimeout,
	}
}

// Option is a function on the options for a ping.
type Option func(opts *Options)

// Count is an option to set ping count.
func Count(count int) Option {
	return func(opts *Options) {
		opts.Count = count
	}
}

// Timeout is an option to set total timeout for ping.
func Timeout(d time.Duration) Option {
	return func(opts *Options) {
		opts.Timeout = d
	}
}

// GetLatency returns the avg ping latency for the addr
// addr can be ip string or domain name
func GetLatency(addr string, options ...Option) (Latency, error) {
	pinger, err := pping.NewPinger(addr)
	if err != nil {
		return Latency{}, err
	}

	opts := DefaultOptions()
	for _, opt := range options {
		opt(&opts)
	}

	pinger.Count = opts.Count
	pinger.Timeout = opts.Timeout
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
