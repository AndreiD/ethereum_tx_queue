package main

import (
	"context"
	"log"
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
	log.Printf("Ethereum Tx Queue: version 0.1")
	log.Printf("Server Port %s", os.Getenv("SERVER_PORT"))
	log.Printf("RPC URL %s", os.Getenv("RPC_URL"))
	log.Printf("Max Queue Size %s", os.Getenv("MAX_QUEUE_SIZE"))
	log.Printf("Errors Retry Count %s", os.Getenv("RETRY_COUNT"))
	log.Printf("Errors Retry Delay %s seconds", os.Getenv("RETRY_DELAY_MS"))
	log.Printf("Local Persistance %s", os.Getenv("LOCAL_DB_PATH"))

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
			log.Fatalf("listen failed: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("server shutdown initiated")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown:", err)
	}
	log.Println("Server Exiting. Bye!")
}
