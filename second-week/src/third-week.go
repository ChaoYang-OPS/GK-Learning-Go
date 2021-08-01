package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func startHttpServer(service *http.Server) error {
	http.HandleFunc("/hello", fmt.Println("......"))
	service, err := service.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	group, errorctx := errgroup.WithContext(ctx)
	//http server
	service := &http.Server{Addr: ":8081"}

	group.Go(func() error {
		return startHttpServer(service)
	})

	group.Go(func() error {
		<-errorctx.Done()
		return service.Shutdown(errorctx)
	})

	chanel := make(chan os.Signal, 1)
	signal.Notify(chanel)

	group.Go(func() error {
		for {
			select {
			case <-errorctx.Done():
				return errorctx.Err()
			case <-chanel:
				cancel()
			}
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		log.Fatal(err)
	}
}
