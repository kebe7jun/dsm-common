package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShaStringBlocks(t *testing.T) {
	res := ShaStringBlocks("a", "b", "c")
	assert.Equal(t, "205830CA-5B23-BBE3-9AB5-10CFDDC1DFF2", res)
	res = ShaStringBlocks("205830ca", "b")
	assert.Equal(t, "C950EF0D-40F6-B9D1-F1B5-D59048B57D0E", res)
}
