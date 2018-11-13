package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/aikuma0130/goWeb/meander"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	meander.APIKey = os.Getenv("GOOGLE_API_KEY")
	if meander.APIKey == "" {
		log.Println("Please Set GOOGLE_API_KEY")
		return
	}
	http.HandleFunc("/journeys", func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, meander.Journeys)
	})
	http.HandleFunc("/recommendations", func(w http.ResponseWriter, r *http.Request) {
		q := &meander.Query{
			Journey: strings.Split(r.URL.Query().Get("journey"), "|"),
		}
		q.Lat, _ = strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
		q.Lng, _ = strconv.ParseFloat(r.URL.Query().Get("lng"), 64)
		q.Radius, _ = strconv.Atoi(r.URL.Query().Get("radius"))
		q.CostRangeStr = r.URL.Query().Get("cost")
		places := q.Run()
		respond(w, r, places)
	})
	http.ListenAndServe(":18080", http.DefaultServeMux)
}

func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	publicData := make([]interface{}, len(data))
	for i, d := range data {
		publicData[i] = meander.Public(d)
	}
	return json.NewEncoder(w).Encode(publicData)
}

func cors(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		f(w, r)
	}
}
