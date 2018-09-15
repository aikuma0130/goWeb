package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/stretchr/gomniauth"

	"github.com/aikuma0130/goWeb/trace"
	"github.com/stretchr/gomniauth/providers/google"
)

type templateHandler struct {
	once     sync.Once
	filename string
	temp1    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.temp1 = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.temp1.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8888", "アプリケーションのアドレス")
	var securityKey = os.Getenv("SECURITY_KEY")
	var githubClientID = os.Getenv("GITHUB_CLIENT_ID")
	var githubSecrets = os.Getenv("GITHUB_SECRETS")
	var googleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	var googleSecrets = os.Getenv("GOOGLE_SECRETS")
	flag.Parse()

	gomniauth.SetSecurityKey(securityKey)
	gomniauth.WithProviders(
		//facebook.New("クライアントID", "秘密の値", "http://localhost:8080/auth/callback/facebook"),
		github.New(githubClientID, githubSecrets, "http://localhost:8080/auth/callback/github"),
		google.New(googleClientID, googleSecrets, "http://localhost:8080/auth/callback/google"),
	)

	r := newRoom()
	r.tracer = trace.New(os.Stdout)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)

	go r.run()

	// Webサーバーを開始します
	log.Println("Webサーバーを開始します。ポート：", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
