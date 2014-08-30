package main

import (
  "fmt"
  "os"
  "time"
  "github.com/gphat/defs"
  "github.com/gphat/system"
)

type HostInfo struct {
  hostname string
}

func main() {

  plugins := map[string]func() defs.Response{ "poop": system.GetMetrics }

  ticker := time.NewTicker(time.Millisecond * 1000)
  go func() {
    for t := range ticker.C {
      for plugin_name, f := range plugins {
        hostname, err := os.Hostname()
        if err != nil {
          fmt.Printf("Shit, error: %v\n", err)
        }
        var hi = HostInfo{hostname: hostname}

        fmt.Printf("Running: %v\n", plugin_name)
        resp := f()
        fmt.Println(resp.Metrics[0])
        fmt.Printf("Hello, from %v\n", hi.hostname)
        fmt.Println("Ticker at", t)
      }
    }
  }()

  time.Sleep(time.Millisecond * 5000)
  ticker.Stop()
  fmt.Println("Ticker stopped")
}
