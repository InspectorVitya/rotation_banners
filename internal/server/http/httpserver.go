package httpserver

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/inspectorvitya/rotation_banners/internal/app"
	"github.com/inspectorvitya/rotation_banners/internal/configuration"
)

type Server struct {
	httpServer *http.Server
	router     *mux.Router
	App        *app.App
}

func New(cfg *configuration.Config, app *app.App) *Server {
	router := mux.NewRouter()

	server := &Server{
		httpServer: &http.Server{
			Addr:         net.JoinHostPort(cfg.ServerHTTP.Host, cfg.ServerHTTP.Port),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
			Handler:      router,
		},
		router: router,
		App:    app,
	}

	return server
}

func (s *Server) Start() error {
	fmt.Println(s.httpServer.Addr)
	s.router.HandleFunc("/banner", s.addBannerInSlot).Methods(http.MethodPost)
	s.App.Logger.Info("start http server...")
	return s.httpServer.ListenAndServe()
}

func (s *Server) Close(ctx context.Context) error {
	s.App.Logger.Info("http server is closing gracefully...")
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("failed to stop http server: %w", err)
	} else {
		s.App.Logger.Info("http server is closed gracefully...")
	}
	return nil
}
