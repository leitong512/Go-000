package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		return App(ctx)
	})

	group.Go(func() error {
		return Debug(ctx)
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	group.Go(func() error {
		for {
			select {
			case s := <-c:
				switch s {
				case syscall.SIGINT, syscall.SIGTERM:
					return errors.New("Close by define signal " + s.String())
				default:
					return errors.New("Close by other signal " + s.String())
				}
			case <-ctx.Done():
				return nil
			}
		}
	})

	if err := group.Wait(); err != nil {
		log.Printf("server error : %s\n", err)
	}
	time.Sleep(2 * time.Second)
	log.Println("exit")
}

func App(ctx context.Context) error {
	m := http.NewServeMux()
	s := &http.Server{
		Addr:    ":6666",
		Handler: m,
	}
	m.HandleFunc("/App", HandleApp)
	go func() {
		<-ctx.Done()
		sdCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		s.Shutdown(sdCtx)
		log.Println("App ShutDown")
	}()
	log.Println("App is staring...")
	return s.ListenAndServe()
}
func HandleApp(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Do HomeWork ...")
}
func HandleDebug(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Debug...")
}
func Debug(ctx context.Context) error {
	m := http.NewServeMux()
	s := &http.Server{
		Addr:    ":6667",
		Handler: m,
	}
	m.HandleFunc("/Debug", HandleDebug)
	go func() {
		<-ctx.Done()
		sdCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		s.Shutdown(sdCtx)
		log.Println("Debug ShutDown")
	}()
	log.Println("Debug is staring...")
	return s.ListenAndServe()

}
