package main

import (
	"fmt"
	"net/http"

	"github.com/serdyanuk/cicd-jenkins/internal/repo/inmemory"
)

func main() {
	searchHistory := inmemory.NewSearchHistoryRepo()

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		if q != "" {
			searchHistory.Add(q)
		}

		history := searchHistory.List()
		for i := len(history) - 1; i >= 0; i-- {
			fmt.Fprintf(w, "#%d. %s\n", i+1, history[i])
		}

		if len(history) == 0 {
			fmt.Fprintf(w, "History is empty")
		}
	})
	http.ListenAndServe(":4000", nil)
}
