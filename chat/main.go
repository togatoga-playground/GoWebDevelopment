package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"mycode/trace"
	"os"
	"github.com/stretchr/gomniauth"
	"io/ioutil"
	"encoding/json"
	"fmt"
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
	t.templ.Execute(w, nil)
}

type authConfig struct {
	Service  string `json:"service"`
	ClientId string `json:"client_id"`
	Secret   string `json:"secret"`
}

func getAuthConfigs() map[string]authConfig {
	bytes, err := ioutil.ReadFile("config/auth.json")
	if err != nil {
		log.Fatal(err)
	}
	var authConfigs []authConfig
	if err := json.Unmarshal(bytes, &authConfigs); err != nil {
		log.Fatal(err)
	}
	res := map[string]authConfig{}
	for _, v := range authConfigs {
		res[v.Service] = v
	}
	return res
}

func main() {
	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	authConfigs := getAuthConfigs()
	gomniauth.SetSecurityKey("togatogatogatogatogatoga")
	gomniauth.WithProviders(
		google.New(authConfigs["google"].ClientId, authConfigs["google"].Secret, "http://localhost:8080/auth/callback/google"),
	)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	go r.run()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
