package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo"
)

func main() {
	// Setup
	e := echo.New()

	e.GET("/slow", func(ctx echo.Context) error {
		time.Sleep(5 * time.Second)
		return ctx.String(http.StatusOK, "retornou depois de 5 segundos...")
	})

	e.GET("/fast", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Ok")
	})

	go func() {
		if err := e.Start(":5000"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Erro ao iniciar o servidor... %s", err.Error())
		}
	}()

	fmt.Println("Servidor iniciado com sucesso!")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	fmt.Println("Iniciando shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Erro ao realizar o graceful shutdown %s", err.Error())
	}

	fmt.Println("Servidor finalizado com sucesso!")

}
