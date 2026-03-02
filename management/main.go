package main

import (
	"log"
	"net/http"
	// nolint:gosec
	_ "net/http/pprof"
	"os"

	"github.com/netbirdio/netbird/management/cmd"
)

func pprofAddr() string {
	listenAddr := os.Getenv("NB_PPROF_ADDR")
	if listenAddr == "" {
		return "localhost:6060"
	}

	return listenAddr
}

func main() {
	go func() {
        pprof_addr := pprofAddr()
		log.Println("Starting pprof server on", pprof_addr)
		log.Println(http.ListenAndServe(pprof_addr, nil))
	}()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
