package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cybozu-go/log"
)

func main() {
	addr := fmt.Sprintf("http://localhost:%s/health", os.Getenv("PORT"))
	_, err := http.Get(addr)
	if err != nil {
		log.ErrorExit(err)
	}
}
