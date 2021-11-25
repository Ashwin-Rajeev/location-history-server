package api

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ReneKroon/ttlcache"
)

// App is the top level struct
type App struct {
	Port  string
	Cache *ttlcache.Cache
}

func NewApp(port string, exp string) *App {
	dur, err := time.ParseDuration(exp)
	if err != nil {
		log.Fatal(err)
	}
	cache := ttlcache.NewCache()
	cache.SetTTL(time.Duration(dur))
	return &App{
		Port:  port,
		Cache: cache,
	}
}

// Serve starts a service backed by an http.Server using default options.
func (a *App) Serve(h http.Handler) {
	if !strings.HasPrefix(a.Port, ":") {
		a.Port = ":" + a.Port
	}
	server := http.Server{
		Addr:         a.Port,
		Handler:      h,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	ServeWithHTTPServer(&server)
}

func ServeWithHTTPServer(hs *http.Server) {
	go func() {
		log.Printf("listening on port %s...\n", hs.Addr)
		err := hs.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal("failed to start server: ", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
}

func (a *App) Start() {
	a.Serve(a.InitRouter())
}
