package resolve

import (
	"log"
	"testing"
)

func TestIPApi(t *testing.T) {
	domain := "qq.com"
	ips, err := GetIPs(domain)
	log.Printf("ips: %v, err: %v", ips, err)
}

func TestLocation(t *testing.T) {
	ps := IPSign{
		IP:   "111.161.64.40",
		Sign: "98c8b46aba964648c861507ee1b84515",
	}
	loc := GetLocation(ps)
	log.Printf("location: %v", loc)
}
