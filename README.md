# gixel.Pixel

One transparent Gif pixel as Go string (8-bit sRGB). 35 bytes of pure beauty.


## Usage example

```go
package main

import (
	"fmt"
	"github.com/zoonman/gixel"
	"log"
	"net/http"
)

// counter of views, I hope 9,223,372,036,854,775,807 is enough for you
var views uint64

// serve content
func handleConnections(w http.ResponseWriter, r *http.Request) {
	views++
	s := ""
	if views > 1 {
		s = "s"
	}
	fmt.Printf("\rYou have %d view%s", views, s)
	w.Write([]byte(gixel.Pixel))
}

// run it
func main() {
	views = 0
	http.HandleFunc("/", handleConnections)
	log.Print("Listening port: 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
```


## Local iMac benchmarks

```
ab -n 10000 -c 64 http://localhost:8000/

Server Hostname:        localhost
Server Port:            8000

Document Path:          /
Document Length:        35 bytes

Concurrency Level:      64
Time taken for tests:   5.050 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      1360000 bytes
HTML transferred:       350000 bytes
Requests per second:    1980.36 [#/sec] (mean)
Time per request:       32.317 [ms] (mean)
Time per request:       0.505 [ms] (mean, across all concurrent requests)
Transfer rate:          263.02 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        1   18   5.2     17      58
Processing:     1   14   5.6     13      59
Waiting:        1   14   5.6     13      58
Total:         11   32   8.0     31      71

Percentage of the requests served within a certain time (ms)
  50%     31
  66%     33
  75%     36
  80%     37
  90%     39
  95%     44
  98%     59
  99%     67
 100%     71 (longest request)
```