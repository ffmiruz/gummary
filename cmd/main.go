package main

import (
	"encoding/json"
	"github.com/ffmiyo/gummary"
	"log"
	"net/http"
	"strings"
)

func main() {
	log.Println("Started...")
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		url := query.Get("q")

		sel := "p"
		rawText, err := gummary.Scrape(sel, url)
		if err != nil {
			log.Println(err)
			return
		}
		ranked := gummary.RankText(rawText)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"texts": strings.Join(ranked, " "),
		})
	})
	http.ListenAndServe(":8080", nil)
}
