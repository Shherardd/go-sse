package main

import (
	"net/http"
	"nitfy/handlers"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	r := http.NewServeMux()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGKILL)
	sig := <-sigs
	println(sig.String())

	handlers.InitRoures(r)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}

}
