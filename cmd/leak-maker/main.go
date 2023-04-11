// reference: https://faun.pub/how-much-is-too-much-the-linux-oomkiller-and-used-memory-d32186f29c9d
// watch ls -lh leak-maker-file
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	LISTEN_PORT = ":8080"
	FILE_NAME   = "leak-maker-file"
)

func main() {
	memoryTicker := time.NewTicker(time.Millisecond * 5)
	leak := make(map[int][]byte)
	i := 0

	go func() {
		for range memoryTicker.C {
			leak[i] = make([]byte, 1024)
			i++
		}
	}()

	fileTicker := time.NewTicker(time.Millisecond * 5)
	go func() {
		f, err := os.Create(FILE_NAME)
		if err != nil {
			log.Fatalln(err)
		}
		buffer := make([]byte, 1024)
		defer f.Close()
		for range fileTicker.C {
			f.Write(buffer)
			f.Sync()
		}
	}()

	fmt.Printf("To start metrics server: http://localhost%s/metrics\n", LISTEN_PORT)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(LISTEN_PORT, nil)
}
