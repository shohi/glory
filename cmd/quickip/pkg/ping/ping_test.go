package ping

import (
	"log"
	"testing"
)

func TestPing_Reachable(t *testing.T) {
	// err := GetLatency("douban.com")
	lc, err := GetLatency("154.8.131.172")
	log.Printf("latency: %v, err: %v", lc, err)
}

func TestPing_Unreachable(t *testing.T) {
	// err := GetLatency("fackebook.com")
	lc, err := GetLatency("31.13.69.129")

	log.Printf("latency: %v, err: %v", lc, err)
}
