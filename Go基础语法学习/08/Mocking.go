package Mocking

import (
	"bytes"
	"fmt"
	"time"
)

func Countdown(buffer *bytes.Buffer) {
	for i := 3; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(buffer, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(buffer, "Go!")
}
