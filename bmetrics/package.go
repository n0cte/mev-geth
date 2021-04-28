package bmetrics

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	once sync.Once
	mu   sync.Mutex

	rcount int
	rfline string
	file   *os.File
)

func init() {
	once.Do(func() {
		f, err := os.OpenFile("~/Downloads/report.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		f.WriteString("txs;1;2;3;4;5;6;7;8;9;10;11;12;13;14;15;16;17;18;19;20\n")
		file = f
	})

}

func track(count int, elapsed time.Duration) {
	mu.Lock()
	defer mu.Unlock()

	if rcount != count {
		if rfline != "" {
			file.WriteString(fmt.Sprintf("%d;%s\n", rcount, rfline))
		}
		rcount = count
		rfline = fmt.Sprintf("%d", elapsed.Microseconds())
	} else {
		rfline = fmt.Sprintf("%s;%d", rfline, elapsed.Microseconds())
	}
}

func TimeTrack(count int, start time.Time) {
	go track(count, time.Since(start))
}
