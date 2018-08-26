package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	temp1    *template.Template
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
		  <html>
		    <head>
			  <title> チャット </title>
		    </head>
		    <body>
			  チャットしましょう !
		    </body>
		  </html>
		`))
	})

	// Webサーバーを開始します
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
