package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	LoadConfig("config.json")

	go func() {
		for {
			fetchData()
			time.Sleep(time.Duration(AppConfig.FetchIntervalS) * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	bindUrl := AppConfig.Server.IP + ":" + AppConfig.Server.Port
	log.Printf("Starting server on %s", bindUrl)
	log.Fatal(http.ListenAndServe(bindUrl, nil))
}
