package forumv2

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(host, port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           host + ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	signal, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		if err := s.httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()
	<-signal.Done()
	log.Print("Shutting down server")
	if err := s.httpServer.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
	return nil
}
