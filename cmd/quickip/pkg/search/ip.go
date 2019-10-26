package search

import (
	"errors"
	"sort"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/shohi/glory/cmd/quickip/pkg/ping"
	"github.com/shohi/glory/cmd/quickip/pkg/resolve"
)

type QuickIP struct {
	resolve.IPSign
	ping.Latency

	Addr string
}

type QuickIPs []QuickIP

func (q QuickIPs) Len() int {
	return len(q)
}

func (q QuickIPs) Less(i, j int) bool {
	return q[i].Latency.Less(q[j].Latency)
}

func (q QuickIPs) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

// TODO: refactor
type quickIPConfig struct {
	location  bool
	pingCount int
}

func quickIP(domain string, conf quickIPConfig) (QuickIPs, error) {
	info, err := resolve.GetIPs(domain)
	if err != nil {
		log.Debugf("%v - err: %v", domain, err)
		return nil, err
	}

	log.Debugf("%v - ips: %v", domain, len(info.Data))
	if len(info.Data) == 0 {
		return nil, errors.New("no ip")
	}

	// TODO: extract pattern
	type result struct {
		sign  resolve.IPSign
		value ping.Latency
		addr  string
		err   error
	}

	resultCh := make(chan result, len(info.Data))

	var wg sync.WaitGroup
	wg.Add(len(info.Data) + 1)

	for _, v := range info.Data {
		go func(s resolve.IPSign) {
			defer wg.Done()
			lat, err := ping.GetLatency(s.IP, ping.Count(conf.pingCount))
			log.Debugf("ip: %v, latency: %v, err: %v", s.IP, lat, err)

			var addr string
			if conf.location {
				addr = resolve.GetLocation(s)
				log.Debugf("ip: %v, location: %v", s.IP, addr)
			}

			resultCh <- result{
				sign:  s,
				value: lat,
				addr:  addr,
				err:   err,
			}
		}(v)
	}

	// watch result
	qps := make(QuickIPs, 0, len(info.Data))
	go func() {
		defer wg.Done()
		for k := 0; k < len(info.Data); k++ {
			res := <-resultCh
			if res.err != nil {
				continue
			}

			qps = append(qps, QuickIP{
				IPSign:  res.sign,
				Latency: res.value,
				Addr:    res.addr,
			})
		}
	}()
	wg.Wait()

	// sort
	sort.Sort(qps)

	return qps, nil
}
