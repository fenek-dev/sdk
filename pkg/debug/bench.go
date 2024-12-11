package debug

import (
	"fmt"
	"time"
)

func Bench(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s: %s \n", name, time.Since(start).String()) //nolint:nolintlint,forbidigo
	}
}
