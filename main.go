package main

import (
	"context"
	"fmt"

	"net/http"
	"os"
	"os/signal"
	"time"
	"txqueue/database"
	"txqueue/eth"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	fmt.Printf("\nEthereum Tx Queue: version 0.1\n\n")
	fmt.Printf("Server Port %s\n", os.Getenv("SERVER_PORT"))
	fmt.Printf("RPC URL %s\n", os.Getenv("RPC_URL"))
	fmt.Printf("Max Queue Size %s\n", os.Getenv("MAX_QUEUE_SIZE"))
	fmt.Printf("Errors Retry Count %s\n", os.Getenv("RETRY_COUNT"))
	fmt.Printf("Errors Retry Delay %s seconds\n", os.Getenv("RETRY_DELAY_MS"))
	fmt.Printf("Local Persistance %s\n\n", os.Getenv("LOCAL_DB_PATH"))

	//init ETH Client
	eth.InitEthClient(os.Getenv("RPC_URL"))

	//init the db
	database.InitDatabase(os.Getenv("LOCAL_DB_PATH"))

	//consume the txes
	go func() {
		Consume()
	}()

	router = gin.New()
	gin.SetMode(gin.ReleaseMode)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	InitializeRouter()

	server := &http.Server{
		Addr:           "localhost:" + os.Getenv("SERVER_PORT"),
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 10, // 1Mb
	}
	server.SetKeepAlivesEnabled(true)

	// Serve'em
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("server shutdown initiated")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}
	fmt.Println("Server Exiting. Bye!")
}
