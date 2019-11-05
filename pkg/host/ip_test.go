package host

import (
	"log"
	"testing"
)

func TestLocalIP(t *testing.T) {
	ip := LocalIP()
	log.Printf("local ip ===> %v", ip)
}
