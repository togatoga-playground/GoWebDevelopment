package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
	"flag"
	"../trace"
	"os"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}


func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	fClag.Parse()
	gomniauth.SetSecurityKey("セキュリティキー")
	gomniauth.WithProviders(
		facebook.New()
	github.New()
	google.New()
	)
	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/", MustAuth(&templateHandler{filename:"chat.html"}))
	http.Handle("/login", &templateHandler{filename:"login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	go r.run()
	log.Println("Webサーバーを開始します。ポート:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil{
		log.Fatal("ListenAndServe:", err)
	}



}
