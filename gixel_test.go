package gixel_test

import (
	"net/http"
	"log"
	"github.com/zoonman/gixel"
	"testing"
)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(gixel.Pixel))
}

func TestGixelPixel(t *testing.T) {
	http.HandleFunc("/", handleConnections)
	log.Print("Listening port: 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}