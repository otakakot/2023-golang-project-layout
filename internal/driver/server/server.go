package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	*http.Server
}

func NewServer(
	port string,
	handler http.Handler,
) *Server {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: handler,
	}

	return &Server{
		Server: srv,
	}
}

func (srv *Server) Run() {
	slog.Info("start server")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	<-ctx.Done()

	slog.Info("start server shutdown")

	ctx, cansel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cansel()

	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}

	slog.Info("done server shutdown")
}
