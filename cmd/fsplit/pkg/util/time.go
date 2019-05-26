package util

import (
	"io"
	"strings"
	"time"
)

// Cheap integer to fixed-width decimal ASCII. Give a negative width to avoid zero-padding.
// refer `log.go`
func itoa(w io.Writer, i int, wid int) error {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	_, err := w.Write(b[bp:])

	return err
}

// FormatTime returns a textual representation of the time in the layout
// YYYY/mm/dd HH:MM:SS
func FormatTime(m time.Time) string {
	var b strings.Builder

	// YYYY/mm/dd
	_ = itoa(&b, m.Year(), 4)
	b.WriteByte('/')
	_ = itoa(&b, int(m.Month()), 2)
	b.WriteByte('/')
	_ = itoa(&b, m.Day(), 2)

	// Black separator
	b.WriteByte(' ')

	// HH:MM:SS
	_ = itoa(&b, m.Hour(), 2)
	b.WriteByte(':')
	_ = itoa(&b, m.Minute(), 2)
	b.WriteByte(':')
	_ = itoa(&b, m.Second(), 2)

	return b.String()
}
