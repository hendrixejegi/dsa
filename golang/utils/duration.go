package utils

import (
	"fmt"
	"time"
)

func Duration(fn func()) {
	start := time.Now()
	fn()
	fmt.Printf("Runtime: %vµs\n", time.Since(start).Microseconds())
}
