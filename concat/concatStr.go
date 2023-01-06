package concatenation

import (
	"bytes"
	"strings"
)

func ConcatWithBuffer() string {
	var b bytes.Buffer

	for i := 0; i < 1000; i++ {
		b.WriteString(randString())
	}

	return b.String()

}

func ConcatWithJoin() string {
	var strs []string
	for i := 0; i < 1000; i++ {
		strs = append(strs, randString())
	}

	return strings.Join(strs, "")
}

func randString() string {
	return "qwerty-xyz"

}
