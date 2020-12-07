package main

import (
	"context"
	"fmt"
	"github.com/devopsws/learn-pipeline-go/pkg/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main()  {
	sm := http.NewServeMux()
	sm.Handle("/", &handler.HelloWorld{})

	svr := http.Server{
		Addr: ":80",
		Handler: sm,
	}

	go func() {
		err := svr.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChain := make(chan os.Signal)
	signal.Notify(sigChain, os.Interrupt)
	signal.Notify(sigChain, os.Kill)

	sig := <- sigChain
	fmt.Println("going to shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30 * time.Second)

	if err := svr.Shutdown(tc); err != nil {
		log.Fatalf("cannot shutdown http server, %v\n", err)
	}
}
