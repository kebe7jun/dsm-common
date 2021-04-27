package utils

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func ShaStringBlocks(blocks ...string) string {
	text := strings.Join(blocks, ",")
	res := sha256.Sum256([]byte(text))
	return fmt.Sprintf("%X-%X-%X-%X-%X", res[0:4], res[4:6], res[6:8], res[8:10], res[10:16])
}
