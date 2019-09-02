package utils

import (
	"fmt"
	"regexp"
)

func FormatCommas(num float32) string {
	str := fmt.Sprintf("%d", int32(num))
	re := regexp.MustCompile("(\\d+)(\\d{3})")
	for n := ""; n != str; {
		n = str
		str = re.ReplaceAllString(str, "$1,$2")
	}
	return str
}
