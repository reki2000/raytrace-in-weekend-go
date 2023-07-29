package core

import (
	"fmt"
	"os"
)

func info(fmtStr string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, fmtStr, args...)
	fmt.Fprintln(os.Stderr)
}

func debug(fmtStr string, args ...interface{}) {
	// fmt.Fprintf(os.Stderr, fmtStr, args...)
	// fmt.Fprintln(os.Stderr)
}
