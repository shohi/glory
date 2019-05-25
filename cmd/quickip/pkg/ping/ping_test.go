package ping

import (
	"log"
	"testing"

	pping "github.com/sparrc/go-ping"
)

func TestPing(t *testing.T) {
	// pinger, err := pping.NewPinger("www.douban.com")

	// use ip address
	pinger, err := pping.NewPinger("154.8.131.172")
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	pinger.Run()                 // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats

	log.Printf("stat: %#v", stats)
}
