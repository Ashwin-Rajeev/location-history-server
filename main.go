package main

import (
	"location-history-server/internal/api"
	"os"
)

const (
	portEnv     = "HISTORY_SERVER_LISTEN_ADDR"
	defaultPort = "8080"
	expEnv      = "LOCATION_HISTORY_TTL_SECONDS"
	defaultExp  = "50s"
)

func main() {
	port := os.Getenv(portEnv)
	if len(port) == 0 {
		port = defaultPort
	}
	exp := os.Getenv(expEnv)
	if len(exp) == 0 {
		exp = defaultExp
	}
	api.NewApp(port, exp).Start()
}
