package search

import (
	"fmt"
	"os"
	"strings"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/shohi/glory/cmd/quickip/pkg/config"
)

type DomainIP struct {
	Domain string
	IPs    QuickIPs
}

type DomainIPs []DomainIP

type search struct {
	conf config.Config
	dps  DomainIPs
}

func newSearch(conf config.Config, domains []string) *search {
	// TODO: filter duplicated
	dps := make(DomainIPs, 0, len(domains))

	for _, v := range domains {
		dps = append(dps, DomainIP{Domain: v})
	}

	return &search{
		conf: conf,
		dps:  dps,
	}
}

func (s *search) do() {
	if len(s.dps) == 0 {
		log.Info("no domain")
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(s.dps) + 1)

	type result struct {
		index int
		value QuickIPs
		err   error
	}

	resultCh := make(chan result, len(s.dps))

	for k, v := range s.dps {
		// NOTE: use executor pool?
		go func(index int, dp DomainIP) {
			defer wg.Done()
			qps, err := quickIP(v.Domain, s.conf.ShowLocation)
			resultCh <- result{
				index: index,
				value: qps,
				err:   err,
			}

		}(k, v)
	}

	// watch result collection
	go func() {
		defer wg.Done()
		for k := 0; k < len(s.dps); k++ {
			res := <-resultCh
			if res.err != nil {
				continue
			}

			s.dps[res.index].IPs = res.value
		}
	}()

	wg.Wait()

	s.dump()
}

func (s *search) dump() {
	var b strings.Builder

	// first dump domain with no avaibable ip
	for _, v := range s.dps {
		if len(v.IPs) != 0 {
			continue
		}

		b.WriteString("# ")
		b.WriteString(v.Domain)
		b.WriteString("\n")
	}

	// second dump domain with avaibable ips
	for _, v := range s.dps {
		if len(v.IPs) == 0 {
			continue
		}

		for _, p := range v.IPs {
			b.WriteString(p.IP)
			b.WriteString("\t")
			b.WriteString(v.Domain)

			if s.conf.ShowLocation {
				b.WriteString("  # ")
				b.WriteString(p.Addr)
				b.WriteString(" ")
				b.WriteString(fmt.Sprintf("%v", p.Duration))
			}

			b.WriteString("\n")
		}
	}

	w := os.Stdout
	fmt.Fprintf(w, b.String())
}

func Search(conf config.Config, domains []string) {
	s := newSearch(conf, domains)
	s.do()
}
