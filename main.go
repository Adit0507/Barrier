package main

import (
	"fmt"
	"time"
	bar "barrier/barrier"
)

func workandWait(name string, timeToWork int, b*bar.Barrier) {
	start := time.Now()

	for {
		fmt.Println(time.Since(start), name, "is running")
		time.Sleep(time.Duration(timeToWork) * time.Second)
		fmt.Println(time.Since(start), name, "is waiting on barrier")
		b.Wait()
	}
}

func main() {
	barrier := bar.NewBarrier(2)

	go workandWait("SchoolboyQ", 4, barrier)
	go workandWait("Kendrick ", 10, barrier)

	time.Sleep(40 * time.Second)
}
