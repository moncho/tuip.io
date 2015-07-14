package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)
func getIp(r *http.Request) (string, error) {
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded == "" {
		ip, err := net.ResolveTCPAddr("tcp", r.RemoteAddr)
		return ip.IP.String(), err
	} else {
		return forwarded, nil
	}

}

func handler()  http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, err := getIp(r)
		if err == nil {
			fmt.Printf("[%v] Somebody with ip %s wants to know.\n",
				time.Now().Format(time.RFC3339),
				ip)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s\n", ip)
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
