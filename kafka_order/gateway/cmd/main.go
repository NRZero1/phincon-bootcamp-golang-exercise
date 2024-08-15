package main

import (
	"context"
	"fmt"
	"gateway/internal/provider/routes"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	router := gin.New()

	globalRoutesGroup := router.Group("")
	{
		routes.UserRoutes(globalRoutesGroup.Group("/user"), "http://localhost:8080")
		routes.PackageRoutes(globalRoutesGroup.Group("/package"), "http://localhost:8081")
	}
	// protectedGroup := router.Group("")
	// protectedGroup.Use(middleware.Authenticate())
	// {
	// 	routes.TicketOrderRoutes(protectedGroup.Group("/ticket-order"), handler.TicketOrderHandler)
	// }
	server := &http.Server{
		Addr:    "localhost:8090",
		Handler: router,
	}

	go func() {
		fmt.Println("Server is running in port ", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msg(fmt.Sprintf("listen: %s\n", err))
		}
	}()

	quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Info().Msg("Shutting down the server...")

    // Set a timeout for shutdown (for example, 5 seconds).
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := server.Shutdown(ctx); err != nil {
        log.Fatal().Msg(fmt.Sprintf("Server shutdown error: %v", err))
    }
    log.Info().Msg("Server gracefully stopped")
}
