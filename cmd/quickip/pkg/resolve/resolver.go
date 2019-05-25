package resolve

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	SpaceChar = ' '
)

var (
	client = http.Client{}

	ipAPIFormat       = "http://site.ip138.com/domain/read.do?domain=%s&time=%v"
	locationAPIForamt = "http://api.ip138.com/query/?ip=%s&oid=5&mid=5&datatype=jsonp&sign=%s"
)

type IPInfo struct {
	Status bool     `json:"status"`
	Data   []IPSign `json:"data"`
}

type IPSign struct {
	IP   string `json:"ip"`
	Sign string `json:"sign"`
}

type Address [6]string

func (a Address) String() string {
	var b strings.Builder
	var length = len(a)
	for k, v := range a {
		if len(v) == 0 {
			continue
		}
		b.WriteString(v)
		if k != length-1 && len(a[k+1]) > 0 {
			b.WriteByte(SpaceChar)
		}
	}

	return b.String()
}

type Location struct {
	Status  string  `json:"ret"`
	Message string  `json:"msg"`
	Addr    Address `json:"data"` // [country, provice, city, provider, other info 1, other info 2]
}

func getJSON(url string, v interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return err
	}

	return nil
}

func GetIPs(domain string) (*IPInfo, error) {
	tsInMillis := time.Now().UnixNano() / 1e6
	url := fmt.Sprintf(ipAPIFormat, domain, tsInMillis)

	var info IPInfo
	err := getJSON(url, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func GetLocation(p IPSign) string {
	url := fmt.Sprintf(locationAPIForamt, p.IP, p.Sign)

	var loc Location
	err := getJSON(url, &loc)
	if err != nil {
		return ""
	}

	return loc.Addr.String()
}
