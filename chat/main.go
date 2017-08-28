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
	"github.com/stretchr/gomniauth/providers/google"
	"io/ioutil"
	"encoding/json"
	"fmt"
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
	service string `json:"service"`
	clientId string `json:"client_id"`
	secret string `json:"secret"`
}


func getAuthConfigs() []authConfig {
	bytes, err := ioutil.ReadFile("config/auth.json")
	if err != nil {
		log.Fatal(err)
	}
	var authConfigs []authConfig
	if err := json.Unmarshal(bytes, &authConfigs); err != nil {
		log.Fatal(err)
	}
	return authConfigs
}



func main() {
	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	authConfigs := getAuthConfigs()
	fmt.Println(authConfigs)
	gomniauth.SetSecurityKey("togatogatogatogatogatoga")
	

	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename:"login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	go r.run()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
