package handlers

import (
	"board/globals"
	"net/http"
	"strings"
	"time"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		board := r.FormValue("board")
		content := r.FormValue("content")
		timestamp := time.Now().UTC().Format(time.RFC822)

		globals.DB.Exec("INSERT INTO boards (title) VALUES (?) ON CONFLICT IGNORE", board)
		globals.DB.Exec("INSERT INTO posts (content, board, timestamp) VALUES (?, ?, ?)", content, board, timestamp)

		http.Redirect(w, r, strings.Join([]string{"/b/", board}, ""), http.StatusFound)
		return
	}

	Layout(w, r, "post.html", nil)
}
