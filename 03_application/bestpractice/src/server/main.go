package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/cybozu-go/log"
	"github.com/cybozu-go/well"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		io.WriteString(w, "Hello")
	})
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "healthy")
	})

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	s := &well.HTTPServer{
		Server: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}
	log.Info("Start listening", map[string]interface{}{
		log.FnHTTPHost: addr,
	})

	err := s.ListenAndServe()
	if err != nil {
		log.ErrorExit(err)
	}

	err = well.Wait()
	if err != nil && !well.IsSignaled(err) {
		log.ErrorExit(err)
	}
}
