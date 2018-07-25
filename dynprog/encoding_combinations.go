package dynprog

import (
	"fmt"
)

// EncodeInt encode a number
func EncodeInt(n int) []byte {
	s := "abcd"
	b := []byte(s)[3]
	fmt.Println(b)
	return []byte(s)
}
