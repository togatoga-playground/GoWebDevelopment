package main

import (
	"../trace"
	"bufio"
	"encoding/csv"
	"flag"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"
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

	data := map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}
	t.templ.Execute(w, data)
}

func readOauthCsv(fileName string) map[string][]string {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		panic(err.Error())
	}
	reader := csv.NewReader(bufio.NewReader(file))
	result := map[string][]string{}
	header := true
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if header {
			header = false
			continue
		}
		result[record[0]] = append(result[record[0]], record[1])
		result[record[0]] = append(result[record[0]], record[2])
	}
	return result
}
func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse()
	gomniauth.SetSecurityKey("セキュリティキー")
	authSetting := readOauthCsv("auth.csv")
	gomniauth.WithProviders(
		google.New(authSetting["google"][0], authSetting["google"][1], "http://localhost:8080/auth/callback/google"),
	)
	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	go r.run()
	log.Println("Webサーバーを開始します。ポート:", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
