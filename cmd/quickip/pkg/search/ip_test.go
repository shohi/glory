package search

import (
	"log"
	"testing"
)

func TestQuickIP(t *testing.T) {
	qps, err := quickIP("douban.com", quickIPConfig{location: true})
	log.Printf("quick ips: %v, err: %v", qps, err)
}
