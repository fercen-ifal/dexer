package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fercen-ifal/dexer/middlewares"
	"github.com/fercen-ifal/dexer/router/v1"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Carrega as variáveis de ambiente

	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Não foi possível carregar o arquivo .env!")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("Variável PORT não possui valor definido!")
		port = "8080"
	}

	// Inicia o servidor no modo graceful

	server := &http.Server{Addr: "0.0.0.0:" + port, Handler: service()}
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig

		shutdownCtx, cancelShutdownCtx := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancelShutdownCtx()

		go func() {
			<-shutdownCtx.Done()

			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("Saída suave cancelada, forçando saída...")
			}
		}()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// Inicia o servidor

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-serverCtx.Done()
}

func service() http.Handler {
	log.Println("Iniciando servidor...")

	r := echo.New()

	// TODO: Add CORS policy

	r.Use(middleware.RequestID())
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())
	r.Use(middleware.Gzip())
	r.Use(middleware.RemoveTrailingSlash())
	r.Use(middlewares.RequestIDHeader())
	r.Use(middlewares.AppInfo())

	v1Router := r.Group("/api/v1")
	v1.RegisterRoutes(v1Router)

	log.Println("Servidor inicializado.")

	return r
}
