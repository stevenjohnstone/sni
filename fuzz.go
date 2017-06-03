// +build gofuzz

package sni

import (
	"bufio"
	"bytes"
)

func Fuzz(data []byte) int {
	reader := bytes.NewReader(data)
	bufferedReader := bufio.NewReader(reader)
	c := bufferedConn{bufferedReader, nil, nil}

	_, _, err := ServerNameFromConn(c)
	if err != nil {
		return 1
	}
	return 0
}
