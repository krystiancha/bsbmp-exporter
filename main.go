package main

import (
  "net/http"

  log "github.com/Sirupsen/logrus"
  "github.com/prometheus/client_golang/prometheus"
  "github.com/prometheus/client_golang/prometheus/promhttp"

  bsbmp "github.com/david-igou/bsbmp-exporter/services"
  "github.com/david-igou/bsbmp-exporter/collectors"

)

func main() {

  //Create a new instance of the foocollector and 
  //register it with the prometheus client.
  foo := collectors.NewBsbmpCollector(bsbmp.Client{Host: "abcd", Port: "abcd"})
  prometheus.MustRegister(foo)

  //This section will start the HTTP server and expose
  //any metrics on the /metrics endpoint.
  http.Handle("/metrics", promhttp.Handler())
  log.Info("Beginning to serve on port :8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}