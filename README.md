# gixel.Pixel

One transparent Gif pixel as Go string (8-bit sRGB).

```go
import (
	"net/http"
	"log"
	"github.com/zoonman/gixel"
)

// serve content
func handleConnections(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(gixel.Pixel))
}
// run it
func main() {
	http.HandleFunc("/", handleConnections)
	log.Print("Listening port: 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
```

 