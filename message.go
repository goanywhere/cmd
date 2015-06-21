package cmd

import (
	"fmt"
	"time"
)

// Loading provides a simple dotted loading prompt while
// processing task. NOTE that task needs to manually notify
// the prompt once the task is done.
func Loading(done chan bool) {
	var (
		index      = 0
		indicators = []string{"-", "\\", "|", "/"}
	)
	go func() {
		for {
			select {
			case <-done:
				fmt.Print(" \b")
				close(done)
				return

			default:
				fmt.Printf("%s\b", indicators[index])
				if index < (len(indicators) - 1) {
					index++
				} else {
					index = 0
				}
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()
}

func Prompt(message string, values ...interface{}) {
	fmt.Printf(fmt.Sprintf("* %s", message), values...)
}
