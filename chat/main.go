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
	"github.com/stretchr/objx"

	"github.com/aikuma0130/goWeb/trace"
	//"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
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

	data := map[string]interface{}{
		"Host": r.Host,
	}

	if authCookie, err := r.Cookie("auth"); err == nil {
		data["UserData"] = objx.MustFromBase64(authCookie.Value)
	}

	t.temp1.Execute(w, data)
}

func main() {
	var addr = flag.String("addr", ":8888", "アプリケーションのアドレス")
	var securityKey = os.Getenv("SECURITY_KEY")
	//var facebookClientID = os.Getenv("FACEBOOK_CLIENT_ID")
	//var facebookSecrets = os.Getenv("FACEBOOK_SECRETS")
	var githubClientID = os.Getenv("GITHUB_CLIENT_ID")
	var githubSecrets = os.Getenv("GITHUB_SECRETS")
	var googleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	var googleSecrets = os.Getenv("GOOGLE_SECRETS")
	flag.Parse()

	gomniauth.SetSecurityKey(securityKey)
	gomniauth.WithProviders(
		//facebook.New(facebookClientID, facebookSecrets, "http://localhost:8080/auth/callback/facebook"),
		github.New(githubClientID, githubSecrets, "http://localhost:8888/auth/callback/github"),
		google.New(googleClientID, googleSecrets, "http://localhost:8888/auth/callback/google"),
	)

	//r := newRoom(UseAuthAvatar)
	r := newRoom(UseGravatarAvatar)
	r.tracer = trace.New(os.Stdout)
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:   "auth",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		})
		w.Header()["Location"] = []string{"/chat"}
		w.WriteHeader(http.StatusTemporaryRedirect)
	})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	http.Handle("/upload", &templateHandler{filename: "upload.html"})

	go r.run()

	// Webサーバーを開始します
	log.Println("Webサーバーを開始します。ポート：", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
