package day4

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Mine(prefix string) int {
	i := 0
	for {
		md5 := md5.New()
		io.WriteString(md5, "bgvyzdsv"+strconv.Itoa(i))
		if strings.HasPrefix(fmt.Sprintf("%x", md5.Sum(nil)), prefix) {
			break
		}
		i++
	}
	return i
}
