package serde

import (
	"bytes"

	"github.com/bsm/redeo/resp"
)

// REdis Serialization Protocol (RESP)

func DeseializeRESP(data []byte) (*resp.Command, error) {
	r := resp.NewRequestReader(bytes.NewReader(data))

	return r.ReadCmd(nil)
}

func SerializeRESP(c *resp.Command) []byte {
	var buf bytes.Buffer
	w := resp.NewRequestWriter(&buf)

	args := make([][]byte, len(c.Args))
	for k, v := range c.Args {
		args[k] = []byte(v)
	}
	w.WriteCmd(c.Name, args...)
	w.Flush()

	return buf.Bytes()
}

func SerializeRawRESP(name string, args []string) []byte {
	var buf bytes.Buffer
	w := resp.NewRequestWriter(&buf)

	w.WriteCmdString(name, args...)
	w.Flush()

	return buf.Bytes()
}
