package fsutil

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFiles(t *testing.T) {
	assert := assert.New(t)

	s := Files{
		"/path/a@part3",
		"/path/a@part2",
		"/path/a@part1",
	}

	sort.Sort(s)

	assert.Equal(s[0], "/path/a@part1")
	assert.Equal(s[1], "/path/a@part2")
	assert.Equal(s[2], "/path/a@part3")
}
