package host

import (
	"net"
)

const defaultIP = "127.0.0.1"

func LocalIP() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return defaultIP
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			ipv4 := ip.To4()
			if ipv4 != nil && ipv4.String() != defaultIP {
				return ipv4.String()
			}
		}
	}

	// If no ip assigned, use the default one
	return "127.0.0.1"
}
