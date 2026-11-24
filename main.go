package main

import (
	"fmt"
	"strconv"
	"strings"
)

// CompareVersions compares dotted version strings numerically per segment:
// -1 if a < b, 0 if equal, 1 if a > b. Missing segments count as zero, so
// "1.0" equals "1.0.0"; leading zeros are ignored ("1.01" equals "1.1").
func CompareVersions(a, b string) int {
	as := strings.Split(a, ".")
	bs := strings.Split(b, ".")
	for i := 0; i < len(as) || i < len(bs); i++ {
		x, y := 0, 0
		if i < len(as) {
			x, _ = strconv.Atoi(as[i])
		}
		if i < len(bs) {
			y, _ = strconv.Atoi(bs[i])
		}
		switch {
		case x < y:
			return -1
		case x > y:
			return 1
		}
	}
	return 0
}

func main() {
	fmt.Println(CompareVersions("1.2", "1.10"))
	fmt.Println(CompareVersions("1.0", "1.0.0"))
}
