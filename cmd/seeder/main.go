package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Channel to catch the interrupt signal
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	// Channel to control the goroutines
	done := make(chan bool)

	// Start goroutines for making HTTP requests
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				// Perform the GET request
				resp, err := http.Get("http://localhost:8080/books")
				if err != nil {
					fmt.Println("Error making request:", err)
					continue
				}

				_, err = io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("Error reading response:", err)
				}
				resp.Body.Close()
			}
		}
	}()

	// Wait for interrupt signal
	<-stopChan
	fmt.Println("Interrupt received, shutting down...")
	close(done)
}
