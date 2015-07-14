package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func handler()  http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, err := net.ResolveTCPAddr("tcp", r.RemoteAddr)
		if err == nil {
			fmt.Printf("[%v] Somebody with ip %s wants to know.\n",
				time.Now().Format(time.RFC3339),
				ip.IP.String())
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s\n", ip.IP.String())

		} else {
			http.Error(w, "Nope", 500)
		}
	})
}

func main() {
	fmt.Println("Listening on port 8000, good things are bound to happen...")
	http.Handle("/", handler())
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
