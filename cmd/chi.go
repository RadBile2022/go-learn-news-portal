package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/spf13/viper"
	"go-learn-news-portal/internal/framework/primary/rest"
	"go-learn-news-portal/internal/framework/primary/rests"
	appmiddleware "go-learn-news-portal/library/v1/middleware"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func ChiStart() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	core := initCore()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(appmiddleware.WithRequestInfo)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Ganti dengan domain spesifik jika perlu
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // 5 menit cache preflight request
	}))

	r.Route("/api", func(r chi.Router) {
		r.Mount("/", rests.NewAuth(core.authCore).Start(ctx))

		r.Route("/admin", func(r chi.Router) {
			r.Use(appmiddleware.AuthMiddleware())

			r.Mount("/users", rests.NewUser(core.userCore).Start(ctx))
			r.Mount("/categories", rests.NewCategory(core.categoryCore).Start(ctx))
			r.Mount("/contents", rests.NewContent(core.contentCore).Start(ctx))
			r.Mount("/files", rest.NewFile(core.fileCore).Start(ctx))
			fmt.Println("Checking routes after mounting:")
			_ = chi.Walk(r, func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
				fmt.Printf("%s %s\n", method, route)
				return nil
			})
		})

		r.Mount("/fe", rests.NewFrontEnd(core.categoryCore, core.contentCore).Start(ctx))
	})

	httpPort := fmt.Sprintf(":%v", viper.GetString("APP_PORT"))
	httpTimeout := time.Duration(30) * time.Second

	httpServer := &http.Server{
		Addr:         httpPort,
		Handler:      http.TimeoutHandler(r, httpTimeout, ""),
		ReadTimeout:  httpTimeout,
		WriteTimeout: httpTimeout + 2,
	}
	idleConnsClosed := make(chan struct{})
	go func() {
		signint := make(chan os.Signal, 1)
		signal.Notify(signint, os.Interrupt)
		<-signint

		if err := httpServer.Shutdown(ctx); err != nil {
			slog.Error("unexpected error during server shutdown", slog.Any("bind_addr", err))
		}
		close(idleConnsClosed)
	}()

	slog.Info("Server listening on", slog.Any("port", httpPort))

	err := httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("Error starting server", slog.Any("error", err))
	}
}
