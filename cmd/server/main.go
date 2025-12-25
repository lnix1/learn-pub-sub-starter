package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/lnix1/learn-pub-sub-starter/internal/pubsub"
	"github.com/lnix1/learn-pub-sub-starter/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")

	const connStr = "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(connStr)
	if err != nil {
		log.Fatal("error establishing connection to localhost")
	}
	defer conn.Close()

	fmt.Println("Peril server successfully started.")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Program running... Press Ctrl+C to stop.")

	go func () {
		channel, err := conn.Channel
		if err != nil {
			log.Fatal("error opening new channel")
		}

		pubsub.PublishJSON(
			channel, 
			routing.ExchangePerilDirect, 
			routing.PauseKey,
			routing.PlayingState{IsPaused: true
		)
	}

	sig := <-sigChan
	fmt.Printf("\nReceived signal: %v. Shutting down...\n", sig)
	
}
