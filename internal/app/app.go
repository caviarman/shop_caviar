package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/caviarman/shop_caviar/internal/config"
	"github.com/caviarman/shop_caviar/internal/entity"
	"github.com/caviarman/shop_caviar/internal/logger"
	"github.com/caviarman/shop_caviar/internal/migrate"
	"github.com/caviarman/shop_caviar/internal/migrations"
	"github.com/caviarman/shop_caviar/internal/repository"
	"github.com/caviarman/shop_caviar/internal/server"
)

const host = "xn--80abwhe0h.online"

func redirectToTls(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fmt.Sprintf("https://%s%s", host, r.RequestURI), http.StatusMovedPermanently)
}

func Run(conf *config.Config) error {
	repo, err := repository.New(conf.DataBase.URL, conf.DataBase.MaxPoolSize)
	if err != nil {
		return fmt.Errorf("repository.New: %w", err)
	}

	defer repo.Close()

	err = migrate.Run(conf.DataBase.URL, migrations.FS)
	if err != nil {
		return fmt.Errorf("migrate.Run: %w", err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	go func() {
		if err := http.ListenAndServe(":80", http.HandlerFunc(redirectToTls)); err != nil {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	r.Get("/api/user", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		users, err := repo.GetUsers(ctx)
		if err != nil {
			logger.Error(err, "repo.GetUsers")

			bz, _ := json.Marshal(err)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)

			w.Write(bz)

			return
		}

		bz, _ := json.Marshal(users)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(bz)
	})

	r.Post("/api/user", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		defer r.Body.Close()

		var user entity.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			logger.Error(err, "json.NewDecoder.Decode user")

			bz, _ := json.Marshal(err)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			w.Write(bz)

			return
		}

		id, err := repo.CreateUser(ctx, &user)
		if err != nil {
			logger.Error(err, "repo.CreateUser")

			bz, _ := json.Marshal(err)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)

			w.Write(bz)

			return
		}

		bz, _ := json.Marshal(map[string]int{"id": id})

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write(bz)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		dir, file := path.Split(r.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			http.ServeFile(w, r, "/web/dist/web/index.html")
		} else {
			http.ServeFile(w, r, "/web/dist/web/"+path.Join(dir, file))
		}
	})

	httpServer := server.New(r, server.Port(conf.Port))

	waitSignal(httpServer)

	return nil
}

func waitSignal(httpServer *server.Server) {
	fmt.Printf("App started on %s!", httpServer.Port())

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		fmt.Println("shutdown signal: " + s.String())
	case err := <-httpServer.Notify():
		fmt.Println(err, "waitSignal - httpServer.Notify")
	}

	err := httpServer.Shutdown()
	if err != nil {
		fmt.Println(err, "waitSignal - httpServer.Shutdown")
	}
}
