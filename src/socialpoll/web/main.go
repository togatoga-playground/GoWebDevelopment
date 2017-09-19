package main

import (
	"flag"
	"net/http"
	"log"
)

func main()  {
	var addr = flag.String("addr", ":8081", "Webサイトのアドレス")
	flag.Parse()
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("public"))))
	log.Println("Webサイトのアドレス:", *addr)
	http.ListenAndServe(*addr, mux)
}
