package proxy

import (
	"fmt"
	"regexp"
	"strings"
)

type Rule struct {
	re     *regexp.Regexp
	target string
}

const (
	SepComma = ","
	SepBar   = "|"
)

// parseRules parses rules from rule string, where rules string is in format of
// `ptn|addr,ptn|addr`
func parseRules(ruleStr string) ([]Rule, error) {
	var ret []Rule
	if len(ruleStr) == 0 {
		return ret, nil
	}

	tokens := strings.Split(ruleStr, SepComma)
	for _, v := range tokens {
		ts := strings.Split(v, SepBar)
		if len(ts) < 2 {
			continue
		}

		ptn := ts[0]
		addr := strings.Join(ts[1:], SepBar)

		re, err := regexp.Compile(fmt.Sprintf(".*%s.*", ptn))
		if err != nil {
			return nil, err
		}
		ret = append(ret, Rule{re: re, target: addr})
	}

	return ret, nil
}
