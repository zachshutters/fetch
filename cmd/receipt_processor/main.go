package main

import (
	"log"
	"net/http"

	"receipt_processor/internal/receipt"
	"receipt_processor/internal/server"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())

	receiptService := receipt.NewReceiptService()
	receiptServer, err := server.New(receiptService)

	if err == nil {

		server.RegisterHandlers(e, receiptServer)

		// Start the server
		log.Println("Starting server on :8080")
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	} else {
		log.Fatalf("Error starting server: %v", err)
	}

}
