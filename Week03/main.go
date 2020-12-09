package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func startServer1(ctx context.Context) error {
	srv := &http.Server{Addr: ":8081"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world\n")
	})

	result := make(chan error)

	go func() {
		result <- srv.ListenAndServe()
	}()

	select {
	case err := <-result:
		fmt.Printf("server1 error %v\n", err)
		return err
	case <-ctx.Done():
		srv.Shutdown(context.TODO()) // what context should I use here?
		fmt.Println("server1 shutdown")
		return nil
	}
}

func startServer2(ctx context.Context) error {
	srv := &http.Server{Addr: ":8082"}

	http.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "debugging...\n")
	})

	result := make(chan error)

	go func() {
		result <- srv.ListenAndServe()
	}()

	select {
	case err := <-result:
		fmt.Printf("server2 error %v\n", err)
		return err
	case <-ctx.Done():
		srv.Shutdown(context.TODO()) // what context should I use here?
		fmt.Println("server2 shutdown")
		return nil
	}
}
func main() {
	group, ctx := errgroup.WithContext(context.Background())
	ctx, cancelFunc := context.WithCancel(ctx)

	group.Go(func() error {
		return startServer1(ctx)
	})
	group.Go(func() error {
		return startServer2(ctx)
	})

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	fmt.Println("interrupted. stop all servers.")
	cancelFunc()
	time.Sleep(10 * time.Second)
}
